package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	kabusapi "github.com/masayoshi4649/KabuStationAPI"
)

// runHTTPServer は、Gin を利用した HTTP サーバーを起動するエントリーポイントです。
//
// 機能:
//   - ルーターを初期化し、標準的なロガーとリカバリーをミドルウェアとして設定する
//   - ルーティングを登録してから指定ポートで待ち受けを開始する
//
// 引数と型:
//   - なし
//
// 返り値と型:
//   - error: サーバー起動やバインドに失敗した場合のエラー
func runHTTPServer() error {
	gin.SetMode(gin.ReleaseMode)
	rt := gin.New()
	rt.Use(gin.Logger(), gin.Recovery())

	// ----------------------------------------
	registerHTTPRoutes(rt)
	log.Printf("HTTPサーバーが %s で待機開始", httpListenAddr)
	return rt.Run(httpListenAddr)
}

// registerHTTPRoutes は、HTML テンプレート、静的ファイル、API エンドポイントのルーティングを Gin エンジンに登録します。
//
// 機能:
//   - テンプレートをロードし、静的コンテンツの配信パスを設定する
//   - インデックス画面と板情報 API のハンドラを紐付ける
//
// 引数と型:
//   - rt *gin.Engine: ルートを登録する Gin エンジン
//
// 返り値と型:
//   - なし
func registerHTTPRoutes(rt *gin.Engine) {
	rt.LoadHTMLGlob("view/*.html")
	rt.Static("/static", "./view")

	rt.GET("/", handleIndexGET)
	rt.GET("/book", handleBookGET)

	rt.POST("/order/cancel", handleOrderCancelPOST)
	rt.POST("/order/close", handleOrderClosePOST)
	rt.POST("/order/open", handleOrderOpenPOST)

}

// handleIndexGET は、先物コードと限月をタイトルとしてテンプレートに渡し、インデックスページを描画します。
//
// 機能:
//   - テンプレートに埋め込むタイトル文字列を生成する
//   - HTTP ステータス 200 で index.html を返却する
//
// 引数と型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値と型:
//   - なし
func handleIndexGET(c *gin.Context) {
	title := fmt.Sprintf("%s（%s）", cfg.Trade.FutureCode, cfg.Trade.DerivMonth)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   title,
		"onetick": cfg.Trade.OneTick,
		"current": current,
	})
}

// handleBookGET は、現在保持している板データをコピーし、JSON としてレスポンスに返却します。
//
// 機能:
//   - 共有メモリ上の板情報を読み取りロックで保護した上で取り出す
//   - スナップショットを JSON として返却する
//
// 引数と型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値と型:
//   - なし
func handleBookGET(c *gin.Context) {
	bookMu.RLock()
	rows := make([]BookRow, len(orderBook))
	copy(rows, orderBook)
	bookMu.RUnlock()
	c.JSON(http.StatusOK, rows)
}

type ReqOrderCancelPOST struct {
	Long  bool `json:"long"`
	Short bool `json:"short"`
}

type ReqOrderClosePOST struct {
	Long       bool `json:"long"`
	Short      bool `json:"short"`
	OnlyProfit bool `json:"only_profit"`
}

type ReqOrderOpenPOST struct {
	Side           string  `json:"side"`
	FrontOrderType int     `json:"front_order_type"`
	StartPrice     float64 `json:"start_price"`
	Interval       float64 `json:"interval"`
	Qty            int     `json:"qty"`
}

// handleOrderCancelPOST は、注文取消用の JSON ペイロードを受信し、現在の注文情報を取得した上で取消処理を実行します。
//
// 機能:
//   - long/short の取消対象を示すペイロードを JSON で受け取り、内容をログへ記録する
//   - KabuStationAPI から注文一覧を取得し、取消処理の前提となる情報を確認する
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - なし
func handleOrderCancelPOST(c *gin.Context) {
	var req ReqOrderCancelPOST
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("handleOrderCancelPOST ペイロードの解析に失敗", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "リクエストボディが不正です",
		})
		return
	}

	// ----------------------------------------
	if req.Long {
		// 注文中データ取得
		code, res, err := kabusapi.GetInfoOrders(kabusapi.ReqGetInfoOrders{
			Symbol:     codeSymbol,
			Side:       "2",
			Cashmargin: "2",
			State:      "3",
		})
		if err != nil {
			log.Println("handleOrderCancelPOST", err)
			c.Error(err)
			c.Abort()
			return
		}
		if code != 200 {
			log.Println("handleOrderCancelPOST", code)
			c.Status(code)
			c.Abort()
			return
		}

		// 注文取消実行
		for _, v := range res {
			b, _ := json.MarshalIndent(v, "", "  ")
			fmt.Println("----------------------------------------")
			fmt.Println(string(b))

			code, res, err := kabusapi.PutOrderCancelorder(kabusapi.ReqPutOrderCancelorder{OrderId: v.ID})
			if err != nil {
				log.Println("PutOrderCancelorder", err)
				c.Error(err)
				c.Abort()
				return
			}
			if code != 200 {
				log.Println("PutOrderCancelorder", code)
				c.Status(code)
				c.Abort()
				return
			}

			fmt.Println("注文取消", res.Result, res.OrderId)
		}
	}

	// ----------------------------------------
	if req.Short {
		// 注文中データ取得
		code, res, err := kabusapi.GetInfoOrders(kabusapi.ReqGetInfoOrders{
			Symbol:     codeSymbol,
			Side:       "1",
			Cashmargin: "2",
			State:      "3",
		})
		if err != nil {
			log.Println("handleOrderCancelPOST", err)
			c.Error(err)
			c.Abort()
			return
		}
		if code != 200 {
			log.Println("handleOrderCancelPOST", code)
			c.Status(code)
			c.Abort()
			return
		}

		// 注文取消実行
		for _, v := range res {
			b, _ := json.MarshalIndent(v, "", "  ")
			fmt.Println("----------------------------------------")
			fmt.Println(string(b))

			code, res, err := kabusapi.PutOrderCancelorder(kabusapi.ReqPutOrderCancelorder{OrderId: v.ID})
			if err != nil {
				log.Println("PutOrderCancelorder", err)
				c.Error(err)
				c.Abort()
				return
			}
			if code != 200 {
				log.Println("PutOrderCancelorder", code)
				c.Status(code)
				c.Abort()
				return
			}

			fmt.Println("注文取消", res.Result, res.OrderId)
		}
	}
	// ----------------------------------------

	c.Status(http.StatusOK)
}

// handleOrderClosePOST は、建玉決済の希望条件を JSON で受信し、現状は受信内容の記録のみを行うハンドラです。
//
// 主な特徴:
//   - long/short/only_profit の 3 つの真偽値を ShouldBindJSON でバインドする
//   - バインドに失敗した場合は 400 を返し、エラーログへ詳細を出力する
//   - 現段階では受信内容をログに残すのみで、外部 API 呼び出しは行わない
//
// 引数:
//   - c *gin.Context: HTTP リクエストコンテキスト
//
// 返り値:
//   - なし
func handleOrderClosePOST(c *gin.Context) {
	var req ReqOrderClosePOST
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("handleOrderClosePOST ペイロードの読み取りに失敗", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "リクエストボディの形式が正しくありません",
		})
		return
	}

	log.Printf("handleOrderClosePOST long=%t short=%t only_profit=%t\n", req.Long, req.Short, req.OnlyProfit)
	c.Status(http.StatusOK)
}

// handleOrderOpenPOST は、即時注文フォームの内容を JSON で受け取り、ログへ記録するハンドラーです。
//
// 主な特徴:
//   - side/front_order_type/start_price/interval/qty の値を ShouldBindJSON でバインドする
//   - side は売り/買いを表し、"1" が売、"2" が買
//   - front_order_type は種別を表し、指値は 20、逆指値は 30
//   - バインドに失敗した場合は 400 を返し、エラー内容をログ出力する
//   - 本実装では受信した値の記録のみを行い、実際の発注処理は行わない
//
// 引数と型:
//   - c *gin.Context: HTTP サーバーのコンテキスト
//
// 返り値と型:
//   - なし
func handleOrderOpenPOST(c *gin.Context) {
	var req ReqOrderOpenPOST
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("handleOrderOpenPOST ペイロードの読込に失敗", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "リクエストボディの形式が正しくありません",
		})
		return
	}

	log.Printf(
		"handleOrderOpenPOST side=%s front_order_type=%d start_price=%.4f interval=%.4f qty=%d\n",
		req.Side,
		req.FrontOrderType,
		req.StartPrice,
		req.Interval,
		req.Qty,
	)

	switch req.Side {
	// 売
	case "1":
		switch req.FrontOrderType {
		case 20: // 指値
			log.Printf("価格プレビュー")

		case 30: // 逆指値
			log.Printf("価格プレビュー")

		}

		// 買
	case "2":
		switch req.FrontOrderType {
		case 20: // 指値
			log.Printf("価格プレビュー")

		case 30: // 逆指値
			log.Printf("価格プレビュー")

		}
	}

	c.Status(http.StatusOK)
}
