package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sync"

	kabusapi "github.com/masayoshi4649/KabuStationAPI"
	"github.com/pelletier/go-toml/v2"
)

var (
	currentQuote kabusapi.Quote
	quoteMu      sync.RWMutex
	orderBook    []BookRow
	bookMu       sync.RWMutex
)

var cfg Config

func main() {
	/*
		起動引数
	*/
	// -c もしくは --config で指定可能に
	var confPath string
	flag.StringVar(&confPath, "c", "app.toml", "path to config file")
	flag.StringVar(&confPath, "config", "app.toml", "path to config file (alias)")
	flag.Parse()

	conf, err := loadConfig(confPath)
	cfg = conf

	if err != nil {
		log.Printf("failed to load config (%s): %v\n", confPath, err)
		log.Println("Enterキーを押してください...")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n') // 改行が来るまでブロック
		os.Exit(1)
	}

	/*
		WebSocket受信
	*/
	kabusapi.SetAPIKey(os.Getenv(cfg.System.EnvName))
	closeWs, err := kabusapi.OpenQuote(updateQuote, updateBook, debugVar)
	if err != nil {
		log.Fatalf("OpenQuote failed: %v", err)
	}
	defer closeWs()
	select {}
}

func debugVar(q kabusapi.Quote) {
	fmt.Println(orderBook)
}

func updateQuote(q kabusapi.Quote) {
	quoteMu.Lock()
	currentQuote = q
	quoteMu.Unlock()
}

// CalcPrice を中心に上下10tick（重複除いて計19行）を生成
func updateBook(q kabusapi.Quote) {

	if cfg.Trade.OneTick <= 0 {
		return
	}

	toTick := func(p float64) int64 { return int64(math.Round(p / cfg.Trade.OneTick)) }
	fromTick := func(t int64) float64 { return float64(t) * cfg.Trade.OneTick }
	cpTick := toTick(q.CalcPrice)

	// 価格→数量
	sells := make(map[int64]float64, 16)
	buys := make(map[int64]float64, 16)
	add := func(m map[int64]float64, price, qty float64) {
		if price <= 0 || qty <= 0 {
			return
		}
		m[toTick(price)] += qty
	}

	// ---- Sell1..10（手書きで列挙）
	add(sells, q.Sell1.Price, q.Sell1.Qty)
	add(sells, q.Sell2.Price, q.Sell2.Qty)
	add(sells, q.Sell3.Price, q.Sell3.Qty)
	add(sells, q.Sell4.Price, q.Sell4.Qty)
	add(sells, q.Sell5.Price, q.Sell5.Qty)
	add(sells, q.Sell6.Price, q.Sell6.Qty)
	add(sells, q.Sell7.Price, q.Sell7.Qty)
	add(sells, q.Sell8.Price, q.Sell8.Qty)
	add(sells, q.Sell9.Price, q.Sell9.Qty)
	add(sells, q.Sell10.Price, q.Sell10.Qty)

	// ---- Buy1..10（手書きで列挙）
	add(buys, q.Buy1.Price, q.Buy1.Qty)
	add(buys, q.Buy2.Price, q.Buy2.Qty)
	add(buys, q.Buy3.Price, q.Buy3.Qty)
	add(buys, q.Buy4.Price, q.Buy4.Qty)
	add(buys, q.Buy5.Price, q.Buy5.Qty)
	add(buys, q.Buy6.Price, q.Buy6.Qty)
	add(buys, q.Buy7.Price, q.Buy7.Qty)
	add(buys, q.Buy8.Price, q.Buy8.Qty)
	add(buys, q.Buy9.Price, q.Buy9.Qty)
	add(buys, q.Buy10.Price, q.Buy10.Qty)

	// ※ Ask/Bid を別途使う場合は必要に応じて（Sell1/Buy1 と二重計上に注意）
	// add(sells, q.AskPrice, q.AskQty)
	// add(buys,  q.BidPrice, q.BidQty)

	// フレーム（高い→安い）を作成
	const radius = 9 // Current を含め上9・下9 = 19行
	rows := make([]BookRow, 0, radius*2+1)
	for t := cpTick + radius; t >= cpTick-radius; t-- {
		rows = append(rows, BookRow{
			Price:   fromTick(t),
			Current: t == cpTick,
			SellQty: sells[t],
			BuyQty:  buys[t],
		})
	}
	bookMu.Lock()
	fmt.Println(rows)
	orderBook = rows
	bookMu.Unlock()
}

func loadConfig(path string) (Config, error) {
	var cfg Config
	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := toml.Unmarshal(b, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
