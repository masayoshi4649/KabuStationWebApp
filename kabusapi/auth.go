package kabusapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// apiKey は X-API-KEY ヘッダーに設定するトークンを保持します。
var apiKey string

// SetAPIKey は X-API-KEY ヘッダー用のトークンを設定します。
//
// 機能:
//   - API 呼び出し時に利用するトークンを上書きして保持する。
//
// 引数:
//   - k (string): 新しく設定する API トークン。
//
// 返り値:
//   - なし。
func SetAPIKey(k string) { apiKey = k }

// APIKey は現在設定されているトークンを返します。
//
// 機能:
//   - SetAPIKey で保持した X-API-KEY 用トークンを参照する。
//
// 引数:
//   - なし。
//
// 返り値:
//   - string: 現在保持している API トークン。未設定の場合は空文字。
func APIKey() string { return apiKey }

// ErrorResponse はエラー時のレスポンスボディです。
type ErrorResponse struct {
	Code    int    `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

// APIError は非2xxのHTTPレスポンスを表すエラーです。
type APIError struct {
	StatusCode int    // HTTPステータスコード
	Code       int    // API固有のエラーコード
	Message    string // APIメッセージ
	Body       string // 生のレスポンスボディ
}

// Error は APIError の内容を整形した文字列として返します。
//
// 機能:
//   - APIError を error インターフェースとして扱えるようにメッセージ化する。
//
// 引数:
//   - なし（レシーバー e に設定された情報を利用）。
//
// 返り値:
//   - string: ステータスコードやエラーコードを含むメッセージ。レシーバーが nil の場合は空文字。
func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return fmt.Sprintf("api error: status=%d code=%d msg=%s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("api error: status=%d body=%s", e.StatusCode, e.Body)
}

// doRequest は内部用の HTTP 呼び出しを行います。必要に応じて X-API-KEY を付与します。
//
// 機能:
//   - HTTP メソッドやパス、クエリ、ボディ、認証要否を受け取り、REST API を実行する。
//
// 引数:
//   - method (string): 呼び出す HTTP メソッド。
//   - path (string): BaseURL と結合する API パス。
//   - query (url.Values): クエリパラメータ。
//   - body ([]byte): リクエストボディ。nil または長さ 0 でボディなし。
//   - needAuth (bool): X-API-KEY ヘッダーを付与するかどうか。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - data ([]byte): レスポンスボディの生データ。
//   - err (error): リクエスト生成・送信時のエラー。
func doRequest(method, path string, query url.Values, body []byte, needAuth bool) (code int, data []byte, err error) {
	u := BaseURL + path
	if query != nil && len(query) > 0 {
		qs := query.Encode()
		if qs != "" {
			if strings.Contains(u, "?") {
				u += "&" + qs
			} else {
				u += "?" + qs
			}
		}
	}
	var r io.Reader
	if len(body) > 0 {
		r = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, u, r)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Accept", "application/json")
	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	if needAuth {
		if apiKey == "" {
			return 0, nil, fmt.Errorf("missing API key: SetAPIKey() が必要です")
		}
		req.Header.Set("X-API-KEY", apiKey)
	}
	resp, err := HTTPClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, b, nil
}

// ReqPostAuthToken は **POST /token** のリクエスト。
//
// ### 概要
// トークン発行
type ReqPostAuthToken struct {
	// APIパスワード
	APIPassword string `json:"APIPassword" validate:"required"`
}

// ResPostAuthToken は **POST /token** のレスポンス。
//
// ### 概要
// トークン発行
type ResPostAuthToken struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	ResultCode int `json:"ResultCode,omitempty"`
	// APIトークン
	Token string `json:"Token,omitempty"`
}

// PostAuthToken は **POST /token** を実行し、API トークンを発行します。
//
// 機能:
//   - kabuステーション API のトークン発行エンドポイントへリクエストを送り、新規トークンを取得する。
//
// 引数:
//   - req (ReqPostAuthToken): API パスワードを含むリクエストボディ。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPostAuthToken): 発行結果コードと取得したトークン。
//   - err (error): 通信エラーまたは APIError。
func PostAuthToken(req ReqPostAuthToken) (code int, res ResPostAuthToken, err error) {
	p := "/token"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := false
	code, data, err := doRequest("POST", p, v, b, needAuth)
	if err != nil {
		return code, res, err
	}
	if code >= 200 && code < 300 {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &res); err != nil {
				return code, res, err
			}
		}
		return code, res, nil
	}
	var apiErr ErrorResponse
	if len(data) > 0 {
		_ = json.Unmarshal(data, &apiErr)
	}
	return code, res, &APIError{StatusCode: code, Code: apiErr.Code, Message: apiErr.Message, Body: string(data)}
}
