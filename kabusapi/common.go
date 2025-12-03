package kabusapi

import (
	"net/http"
	"strings"
	"time"
)

// BaseURL は API のベースURL。
// 変更したい場合は SetBaseURL("http://localhost:18081/kabusapi") のように設定してください。
var BaseURL = "http://localhost:18080/kabusapi"

// HTTPClient は REST 呼び出しに使用されるクライアントです。必要に応じて差し替え可能です。
var HTTPClient = &http.Client{Timeout: 30 * time.Second}

// SetHTTPClient は使用する http.Client を差し替えます。
//
// 機能:
//   - REST 呼び出しに使用する HTTP クライアントを任意の設定に置き換える。
//
// 引数:
//   - c (*http.Client): 使用したいクライアント。nil の場合は既定を維持。
//
// 返り値:
//   - なし。
func SetHTTPClient(c *http.Client) {
	if c != nil {
		HTTPClient = c
	}
}

// SetBaseURL は BaseURL を差し替えます。
//
// 機能:
//   - API の送信先となるベース URL を上書きし、末尾のスラッシュを除去する。
//
// 引数:
//   - u (string): 新しいベース URL。空文字は無視される。
//
// 返り値:
//   - なし。
func SetBaseURL(u string) {
	if u != "" {
		BaseURL = strings.TrimRight(u, "/")
	}
}
