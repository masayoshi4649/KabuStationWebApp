package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"sync"

	kabusapi "github.com/masayoshi4649/KabuStationAPI"
	"github.com/pelletier/go-toml/v2"
)

var (
	orderBook []BookRow
	bookMu    sync.RWMutex
)
var cfg Config

const httpListenAddr = ":8080"

/*
### 機能
設定ファイルを読み込み、HTTPサーバーとWebSocket購読を開始してアプリケーション全体を起動する。

### 引数およびその型
- なし

### 返り値およびその型
- なし
*/
func main() {
	// 設定ファイルの読み込みと検証
	var confPath string
	flag.StringVar(&confPath, "c", "app.toml", "設定ファイルへのパスを指定するフラグ")
	flag.StringVar(&confPath, "config", "app.toml", "設定ファイルへのパスを指定するエイリアス")
	flag.Parse()

	conf, err := loadConfig(confPath)
	cfg = conf
	if err != nil {
		log.Printf("設定ファイルの読込に失敗しました (%s): %v\n", confPath, err)
		log.Println("Enterキーを押すと終了します...")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n') // エラー内容を確認するための一時停止
		os.Exit(1)
	}

	// --------------------
	go func() {
		if err := runHTTPServer(); err != nil {
			log.Printf("HTTPサーバーが停止しました: %v", err)
		}
	}()

	// --------------------
	kabusapi.SetAPIKey(os.Getenv(cfg.System.EnvName))
	closeWs, err := kabusapi.OpenQuote(updateBook, debugVar)
	if err != nil {
		log.Fatalf("OpenQuoteの開始に失敗しました: %v", err)
	}
	defer closeWs()

	select {}
}

/*
### 機能
受信したQuoteと最新の板情報をログに出力し、リアルタイム処理をデバッグしやすくする。

### 引数およびその型
- `q` kabusapi.Quote - APIから受信した板情報。

### 返り値およびその型
- なし
*/
func debugVar(q kabusapi.Quote) {
	log.Println(q)
	log.Println(orderBook)
}

/*
### 機能
受信したQuoteから売買数量をティック単位で集計し、HTTP配信用の板データに更新する。

### 引数およびその型
- `q` kabusapi.Quote - APIから受信した最新の板情報。

### 返り値およびその型
- なし
*/

func updateBook(q kabusapi.Quote) {
	if cfg.Trade.OneTick <= 0 {
		return
	}

	toTick := func(p float64) int64 { return int64(math.Round(p / cfg.Trade.OneTick)) }
	fromTick := func(t int64) float64 { return float64(t) * cfg.Trade.OneTick }
	cpTick := toTick(q.CalcPrice)

	// --------------------
	// 板情報構築用の一時領域
	sells := make(map[int64]float64, 16)
	buys := make(map[int64]float64, 16)
	add := func(m map[int64]float64, price, qty float64) {
		if price <= 0 || qty <= 0 {
			return
		}
		m[toTick(price)] += qty
	}

	// --------------------
	// 売り気配の集計
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

	// --------------------
	// 買い気配の集計
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

	// Ask/Bidを直接加算したい場合は以下を有効化する（Sell1/Buy1との二重計上に注意）
	// add(sells, q.AskPrice, q.AskQty)
	// add(buys, q.BidPrice, q.BidQty)

	// --------------------
	// 表示用の行データへ変換
	const radius = 9 // 現在値を中心に前後9刻みを表示
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
	orderBook = rows
	bookMu.Unlock()
}

/*
### 機能
TOML形式の設定ファイルを読み込み、Config構造体へデシリアライズする。

### 引数およびその型
- `path` string - 読み込む設定ファイルのパス。

### 返り値およびその型
- Config - 読み込んだ設定値。
- error - 読み込みまたは解析に失敗した場合のエラー。
*/

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
