package kabusapi

import (
	"context"
	"errors"
	"log"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

const endpoint = "ws://localhost:18080/kabusapi/websocket"

// Websocket受信データ
type Quote struct {
	AskPrice float64 `json:"AskPrice"`
	AskQty   float64 `json:"AskQty"`
	AskSign  string  `json:"AskSign"`
	BidPrice float64 `json:"BidPrice"`
	BidQty   float64 `json:"BidQty"`
	BidSign  string  `json:"BidSign"`
	Buy1     struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
		Sign  string  `json:"Sign"`
	} `json:"Buy1"`
	Buy10 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy10"`
	Buy2 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy2"`
	Buy3 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy3"`
	Buy4 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy4"`
	Buy5 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy5"`
	Buy6 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy6"`
	Buy7 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy7"`
	Buy8 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy8"`
	Buy9 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Buy9"`
	CalcPrice                float64   `json:"CalcPrice"`
	ChangePreviousClose      float64   `json:"ChangePreviousClose"`
	ChangePreviousClosePer   float64   `json:"ChangePreviousClosePer"`
	ClearingPrice            float64   `json:"ClearingPrice"`
	CurrentPrice             float64   `json:"CurrentPrice"`
	CurrentPriceChangeStatus string    `json:"CurrentPriceChangeStatus"`
	CurrentPriceStatus       float64   `json:"CurrentPriceStatus"`
	CurrentPriceTime         time.Time `json:"CurrentPriceTime"`
	Exchange                 float64   `json:"Exchange"`
	ExchangeName             string    `json:"ExchangeName"`
	HighPrice                float64   `json:"HighPrice"`
	HighPriceTime            time.Time `json:"HighPriceTime"`
	LowPrice                 float64   `json:"LowPrice"`
	LowPriceTime             time.Time `json:"LowPriceTime"`
	OpeningPrice             float64   `json:"OpeningPrice"`
	OpeningPriceTime         time.Time `json:"OpeningPriceTime"`
	PreviousClose            float64   `json:"PreviousClose"`
	PreviousCloseTime        time.Time `json:"PreviousCloseTime"`
	SecurityType             float64   `json:"SecurityType"`
	Sell1                    struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
		Sign  string  `json:"Sign"`
	} `json:"Sell1"`
	Sell10 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell10"`
	Sell2 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell2"`
	Sell3 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell3"`
	Sell4 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell4"`
	Sell5 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell5"`
	Sell6 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell6"`
	Sell7 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell7"`
	Sell8 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell8"`
	Sell9 struct {
		Price float64 `json:"Price"`
		Qty   float64 `json:"Qty"`
	} `json:"Sell9"`
	Symbol            string    `json:"Symbol"`
	SymbolName        string    `json:"SymbolName"`
	TradingValue      float64   `json:"TradingValue"`
	TradingVolume     float64   `json:"TradingVolume"`
	TradingVolumeTime time.Time `json:"TradingVolumeTime"`
	VWAP              float64   `json:"VWAP"`
}

// Handler は既存の定義と合わせる
type Handler func(Quote)

// OpenQuote は WebSocket に接続して板情報を受信し、ハンドラーへ配信します。
//
// 機能:
//   - kabuステーションの WebSocket エンドポイントへ接続し、受信した Quote を逐次ハンドラーへ渡す。
//   - 返された closeFunc を呼び出すことで受信ループと接続を安全に終了できる。
//
// 引数:
//   - handlers (...func(Quote)): 受信した Quote を処理するコールバック。nil は無視される。
//
// 返り値:
//   - closeFunc (func()): 接続と受信を停止させるための関数。
//   - err (error): 初回接続に失敗した場合のエラー。
func OpenQuote(handlers ...func(Quote)) (closeFunc func(), err error) {
	// endpoint と Quote 型は ws.go に定義されている想定（参照: ws.go） :contentReference[oaicite:2]{index=2}
	if endpoint == "" {
		return nil, errors.New("OpenQuote: endpoint is empty")
	}

	// API キーは auth.go の SetAPIKey / APIKey を使って設定されている前提（参照: auth.go） :contentReference[oaicite:3]{index=3}
	key := APIKey()
	if key == "" {
		return nil, errors.New("OpenQuote: API key is empty; call SetAPIKey(...) before OpenQuote")
	}

	// 管理用 context
	ctx, cancel := context.WithCancel(context.Background())

	// 初回ダイヤル（短時間タイムアウト）
	dialCtx, dialCancel := context.WithTimeout(ctx, 5*time.Second)
	defer dialCancel()

	conn, _, dErr := websocket.Dial(dialCtx, endpoint, &websocket.DialOptions{
		HTTPHeader: map[string][]string{
			"X-API-KEY": {key},
		},
	})
	if dErr != nil {
		cancel()
		return nil, dErr
	}

	var wg sync.WaitGroup
	var once sync.Once
	var closing uint32 // 0 == running, 1 == closing initiated

	// safeCall: handler 内の panic を吸収する
	safeCall := func(h func(Quote), q Quote) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("kabusapi: OpenQuote handler panic recovered: %v", r)
			}
		}()
		h(q)
	}

	// read loop
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 終了時は必ずコネクションをノーマルで閉じる（WriteControl は nhooyr が管理するため Close で十分）
		defer func() { _ = conn.Close(websocket.StatusNormalClosure, "") }()

		for {
			var q Quote
			if rErr := wsjson.Read(ctx, conn, &q); rErr != nil {
				// 呼び出し元がキャンセルした（close() が呼ばれた等）
				if errors.Is(rErr, context.Canceled) || errors.Is(rErr, context.DeadlineExceeded) {
					return
				}

				// WebSocket 側の正常クローズ判定
				if s := websocket.CloseStatus(rErr); s == websocket.StatusNormalClosure ||
					s == websocket.StatusGoingAway || s == websocket.StatusNoStatusRcvd {
					return
				}

				// ネットワークが既に閉じられている（race 発生時のノイズ）
				if errors.Is(rErr, net.ErrClosed) || strings.Contains(rErr.Error(), "use of closed network connection") {
					// client 側で close() を発火して race したなら静かに抜ける
					if atomic.LoadUint32(&closing) == 1 {
						return
					}
					// そうでなければログに残して終了（必要なら自動再接続を追加）
					log.Printf("kabusapi: websocket read network closed: %v", rErr)
					return
				}

				// その他のエラーはログを出して終了（再接続ポリシーを入れるならここを拡張）
				log.Printf("kabusapi: websocket read error: %T %v", rErr, rErr)
				return
			}

			// ハンドラを実行（各ハンドラは個別 goroutine で非ブロッキングに）
			for _, h := range handlers {
				if h == nil {
					continue
				}
				go safeCall(h, q)
			}
		}
	}()

	// closeFunc を返す（安全に複数回呼べる）
	closeFn := func() {
		once.Do(func() {
			atomic.StoreUint32(&closing, 1)
			// goroutine 停止を促す
			cancel()
			// conn.Close して Read をアンロックする
			_ = conn.Close(websocket.StatusNormalClosure, "client close")

			// goroutine の終了を待つ
			done := make(chan struct{})
			go func() {
				wg.Wait()
				close(done)
			}()
			select {
			case <-done:
				// 正常終了
			case <-time.After(3 * time.Second):
				log.Println("kabusapi: OpenQuote close timeout waiting for goroutine")
			}
		})
	}

	return closeFn, nil
}
