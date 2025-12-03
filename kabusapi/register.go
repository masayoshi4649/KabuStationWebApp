package kabusapi

import (
	"encoding/json"
	"net/url"
)

// ReqPutRegisterRegister は PUT /register のリクエストボディ。
//   - 用途: PUSH 配信対象の銘柄を登録する
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutRegisterRegister struct {
	Symbols []struct {
		Symbol string `json:"Symbol,omitempty"` // 銘柄コード
		// Exchange は市場コード。
		//   - 1: 東証
		//   - 3: 名証
		//   - 5: 福証
		//   - 6: 札証
		//   - 2: 日通し
		//   - 23: 日中
		//   - 24: 夜間
		Exchange int `json:"Exchange,omitempty"`
	} `json:"Symbols"`
}

// ResPutRegisterRegister は PUT /register のレスポンス。
//   - 返却内容: 登録済み銘柄リスト
type ResPutRegisterRegister struct {
	// 現在登録されている銘柄のリスト
	RegistList []struct {
		Symbol string `json:"Symbol,omitempty"` // 銘柄コード
		// Exchange は市場コード。
		//   - 1: 東証
		//   - 3: 名証
		//   - 5: 福証
		//   - 6: 札証
		//   - 2: 日通し
		//   - 23: 日中
		//   - 24: 夜間
		Exchange int `json:"Exchange,omitempty"`
	} `json:"RegistList,omitempty"`
}

// PutRegisterRegister は **PUT /register** を実行し、PUSH 配信対象の銘柄を登録します。
//
// 機能:
//   - PUSH 配信で受信したい銘柄と市場コードのリストを API 登録銘柄リストに追加する。
//
// 引数:
//   - req (ReqPutRegisterRegister): 登録する銘柄コードと市場コードの配列。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPutRegisterRegister): 登録後の銘柄リスト。
//   - err (error): 通信エラーまたは APIError。
func PutRegisterRegister(req ReqPutRegisterRegister) (code int, res ResPutRegisterRegister, err error) {
	p := "/register"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
	code, data, err := doRequest("PUT", p, v, b, needAuth)
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

// ReqPutRegisterUnregister は PUT /unregister のリクエストボディ。
//   - 用途: PUSH 登録した銘柄を個別に解除する
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutRegisterUnregister struct {
	// 為替銘柄を登録解除する場合の指定:
	//   - Symbol: 通貨A/通貨B（例: EUR/USD）
	//   - Exchange: 300
	Symbols []struct {
		Symbol string `json:"Symbol,omitempty"` // 銘柄コード
		// Exchange は市場コード。
		//   - 1: 東証
		//   - 3: 名証
		//   - 5: 福証
		//   - 6: 札証
		//   - 2: 日通し
		//   - 23: 日中
		//   - 24: 夜間
		//   - 300: 為替（通貨ペア）
		Exchange int `json:"Exchange,omitempty"`
	} `json:"Symbols"`
}

// ResPutRegisterUnregister は PUT /unregister のレスポンス。
//   - 返却内容: 登録解除後の銘柄リスト
type ResPutRegisterUnregister struct {
	// 現在登録されている銘柄のリスト
	RegistList []struct {
		Symbol string `json:"Symbol,omitempty"` // 銘柄コード
		// Exchange は市場コード。
		//   - 1: 東証
		//   - 3: 名証
		//   - 5: 福証
		//   - 6: 札証
		//   - 2: 日通し
		//   - 23: 日中
		//   - 24: 夜間
		//   - 300: 為替（通貨ペア）
		Exchange int `json:"Exchange,omitempty"`
	} `json:"RegistList,omitempty"`
}

// PutRegisterUnregister は **PUT /unregister** を実行し、指定銘柄の PUSH 登録を解除します。
//
// 機能:
//   - API 登録銘柄リストに登録済みの銘柄を指定し、PUSH 配信対象から削除する。
//
// 引数:
//   - req (ReqPutRegisterUnregister): 解除対象の銘柄コードと市場コードのリスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPutRegisterUnregister): 解除後の登録銘柄リスト。
//   - err (error): 通信エラーまたは APIError。
func PutRegisterUnregister(req ReqPutRegisterUnregister) (code int, res ResPutRegisterUnregister, err error) {
	p := "/unregister"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
	code, data, err := doRequest("PUT", p, v, b, needAuth)
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

// ReqPutRegisterUnregisterAll は PUT /unregister/all のリクエストボディ。
//   - 用途: 登録済み銘柄を全件解除する
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutRegisterUnregisterAll struct {
}

// ResPutRegisterUnregisterAll は PUT /unregister/all のレスポンス。
//   - 返却内容: 全解除後の登録銘柄リスト
type ResPutRegisterUnregisterAll struct {
	// 現在登録されている銘柄のリスト。
	// - 全解除が成功した場合は空。
	// - エラーがある場合は解除前の登録銘柄を返す。
	RegistList []interface{}
}

// PutRegisterUnregisterAll は **PUT /unregister/all** を実行し、登録銘柄を一括解除します。
//
// 機能:
//   - API 登録銘柄リストに登録されている全銘柄を PUSH 配信対象からまとめて削除する。
//
// 引数:
//   - req (ReqPutRegisterUnregisterAll): 全解除を指示する空のリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPutRegisterUnregisterAll): 解除後の登録銘柄リスト。
//   - err (error): 通信エラーまたは APIError。
func PutRegisterUnregisterAll(req ReqPutRegisterUnregisterAll) (code int, res ResPutRegisterUnregisterAll, err error) {
	p := "/unregister/all"
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("PUT", p, v, b, needAuth)
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
