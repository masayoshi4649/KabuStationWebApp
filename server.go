package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	rt.GET("/", handleIndexGET)
	rt.Static("/static", "./view")
	rt.GET("/book", handleBookGET)
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
		"title": title,
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
