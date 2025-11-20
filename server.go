package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
### 機能
Ginルーターを初期化してHTTPサーバーを起動する。

### 引数およびその型
- なし

### 返り値およびその型
- error - サーバーの起動や稼働中に失敗した場合のエラー。
*/

func runHTTPServer() error {
	gin.SetMode(gin.ReleaseMode)
	rt := gin.New()
	rt.Use(gin.Logger(), gin.Recovery())

	// --------------------
	registerHTTPRoutes(rt)
	log.Printf("HTTPサーバーを %s で待ち受け開始", httpListenAddr)
	return rt.Run(httpListenAddr)
}

/*
### 機能
受け取ったGinエンジンにHTML/静的ファイル/板情報APIの各エンドポイントを登録する。

### 引数およびその型
- `rt` *gin.Engine - ハンドラを登録するGinエンジン。

### 返り値およびその型
- なし
*/

func registerHTTPRoutes(rt *gin.Engine) {
	rt.GET("/", handleIndexGET)
	rt.Static("/static", "./view")
	rt.GET("/book", handleBookGET)
}

/*
### 機能
トップページとなるHTMLファイルをレスポンスとして送信する。

### 引数およびその型
- `c` *gin.Context - クライアントリクエストを表すGinコンテキスト。

### 返り値およびその型
- なし
*/

func handleIndexGET(c *gin.Context) {
	c.File("./view/index.html")
}

/*
### 機能
現在保持している板情報をコピーし、JSONで返却する。

### 引数およびその型
- `c` *gin.Context - クライアントリクエストを表すGinコンテキスト。

### 返り値およびその型
- なし
*/

func handleBookGET(c *gin.Context) {
	bookMu.RLock()
	rows := make([]BookRow, len(orderBook))
	copy(rows, orderBook)
	bookMu.RUnlock()
	c.JSON(http.StatusOK, rows)
}
