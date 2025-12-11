package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"sync"

	"github.com/masayoshi4649/KabuStationWebApp/kabusapi"
	"github.com/pelletier/go-toml/v2"
)

var (
	current   float64
	orderBook []BookRow
	bookMu    sync.RWMutex
)
var cfg Config
var codeSymbol string

const httpListenAddr = ":8080"

// main は、設定の読み込みと HTTP サーバー、板情報購読の起動を行うアプリケーションのエントリーポイントです。
//
// 機能:
//   - コマンドライン引数から設定ファイルパスを受け取る
//   - 設定の読み込みに失敗した場合はエラーメッセージを表示して終了を待機する
//   - HTTP サーバーと WebSocket 購読を並行して起動する
//
// 引数と型:
//   - なし
//
// 返り値と型:
//   - なし
func main() {
	orderBook = []BookRow{
		{Price: 0, Current: true, SellQty: 0, BuyQty: 0},
	}

	// 設定ファイルパスの取得と初期値設定
	var confPath string
	flag.StringVar(&confPath, "c", "app.toml", "設定ファイルへのパスを指定するフラグ")
	flag.StringVar(&confPath, "config", "app.toml", "設定ファイルへのパスを指定するエイリアス")

	var apikey string
	flag.StringVar(&apikey, "k", "", "APIKEYを指定するフラグ")
	flag.StringVar(&apikey, "key", "", "APIKEYを指定するエイリアス")

	flag.Parse()

	conf, err := loadConfig(confPath)
	cfg = conf
	if err != nil {
		log.Printf("設定ファイルの読み込みに失敗しました (%s): %v\n", confPath, err)
		log.Println("Enterキーを押すと終了します...")
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n') // エラー内容を確認するために一時停止
		os.Exit(1)
	}

	// ----------------------------------------
	go func() {
		if err := runHTTPServer(); err != nil {
			log.Printf("HTTPサーバーが終了しました: %v", err)
		}
	}()

	// ----------------------------------------
	{
		kabusapi.SetAPIKey(apikey)

		// 既存登録を全解除
		code, _, err := kabusapi.PutRegisterUnregisterAll(kabusapi.ReqPutRegisterUnregisterAll{})
		if code != 200 || err != nil {
			log.Println("PutRegisterUnregisterAll", code, err)
			return
		}

		// 対象銘柄を取得
		code, res, err := kabusapi.GetInfoSymbolnameFuture(kabusapi.ReqGetInfoSymbolnameFuture{
			FutureCode: cfg.Trade.FutureCode,
			DerivMonth: cfg.Trade.DerivMonth,
		})
		if code != 200 || err != nil {
			log.Println("GetInfoSymbolnameFuture", code, err)
			return
		}
		codeSymbol = res.Symbol

		// 板情報の登録
		code, _, err = kabusapi.PutRegisterRegister(kabusapi.ReqPutRegisterRegister{
			Symbols: []struct {
				Symbol   string `json:"Symbol,omitempty"`
				Exchange int    `json:"Exchange,omitempty"`
			}{
				{
					Symbol: codeSymbol,
					// Exchange: cfg.Trade.Exchange,
					Exchange: 2, // ここだけハードコーディング
				},
			},
		})
		if code != 200 || err != nil {
			log.Println("PutRegisterRegister", code, err)
			return
		}

		closeWs, err := kabusapi.OpenQuote(updateBook)
		if err != nil {
			log.Fatalf("OpenQuote の開始に失敗しました: %v", err)
		}
		defer closeWs()

		select {}
	}
}

// debugVar は、受信した Quote と現在保持している板データをログに出力し、動作確認に用いるデバッグ用の関数です。
//
// 機能:
//   - 引数で受け取った Quote 全体をログ出力する
//   - サーバーが保持している板データのスナップショットをログ出力する
//
// 引数と型:
//   - q kabusapi.Quote: API から受信した板情報
//
// 返り値と型:
//   - なし
func debugVar(q kabusapi.Quote) {
	log.Println(q)
	log.Println(orderBook)
}

// updateBook は、受信した Quote をティック単位に集約し、HTTP レスポンス用の板データを最新化します。
//
// 機能:
//   - 現在値をティックに丸めて中央ティックを算出する
//   - 売買それぞれの気配数量をティックごとに合算する
//   - 表示範囲（現在値の前後一定幅）の行データを生成し共有メモリに格納する
//
// 引数と型:
//   - q kabusapi.Quote: API から受信した板情報
//
// 返り値と型:
//   - なし
func updateBook(q kabusapi.Quote) {
	if cfg.Trade.OneTick <= 0 {
		return
	}

	current = q.CurrentPrice

	toTick := func(p float64) int64 { return int64(math.Round(p / cfg.Trade.OneTick)) }
	fromTick := func(t int64) float64 { return float64(t) * cfg.Trade.OneTick }
	cpTick := toTick(q.CalcPrice)

	// ----------------------------------------
	// 板表示用の一時集計領域
	sells := make(map[int64]float64, 16)
	buys := make(map[int64]float64, 16)
	add := func(m map[int64]float64, price, qty float64) {
		if price <= 0 || qty <= 0 {
			return
		}
		m[toTick(price)] += qty
	}

	// ----------------------------------------
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

	// ----------------------------------------
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

	// Ask/Bid を同時に加算する場合は以下を利用（Sell1/Buy1 との重複に注意）
	// add(sells, q.AskPrice, q.AskQty)
	// add(buys, q.BidPrice, q.BidQty)

	// ----------------------------------------
	// 表示用の行データへ変換
	const radius = 9 // 現在値を中心に前後 9 本を表示
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

// loadConfig は、TOML 形式の設定ファイルを読み込み、Config 構造体へデシリアライズします。
//
// 機能:
//   - 指定パスのファイルを読み取り、TOML をパースする
//   - パース結果を Config にマッピングして返す
//
// 引数と型:
//   - path string: 読み込む設定ファイルのパス
//
// 返り値と型:
//   - Config: 読み込んだ設定値
//   - error: ファイル読み込みやパースに失敗した場合のエラー
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
