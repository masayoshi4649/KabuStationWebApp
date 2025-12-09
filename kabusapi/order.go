package kabusapi

import (
	"encoding/json"
	"net/url"
)

// ReqPostOrderSendorderFuture は **POST /sendorder/future** のリクエストボディ。
//
// 概要:
//   - 先物銘柄の新規・返済注文を発注する。
//   - X-API-KEY ヘッダーによる認証が必須。
//   - 同一銘柄での同時発注は 5 件程度までを推奨。
type ReqPostOrderSendorderFuture struct {
	// Symbol は先物銘柄コード。
	// 取引最終日に DerivMonth=0（直近限月）を指定した場合は最終日の限月コードが返るため、日中引けで取引終了となる点に注意。
	Symbol string `json:"Symbol" validate:"required"`
	// Exchange は市場コード。2:日通し、23:日中、24:夜間、32:SOR日通し、33:SOR日中、34:SOR夜間。
	// SOR は日経225先物（ラージ/ミニ）、TOPIX（ラージ/ミニ）、東証マザーズ指数、JPX日経400、NYダウなど一部銘柄のみ有効。
	Exchange int `json:"Exchange" validate:"required"`
	// TradeType は取引区分。1:新規、2:返済。
	TradeType int `json:"TradeType" validate:"required"`
	// TimeInForce は有効期間条件。1:FAS、2:FAK、3:FOK。
	// FrontOrderType および市場コードとの組み合わせに制限があるため、必要に応じて仕様書の対応表を参照する。
	TimeInForce int `json:"TimeInForce" validate:"required"`
	// Side は売買区分。1:売、2:買。
	Side string `json:"Side" validate:"required"`
	// Qty は注文数量。
	Qty int `json:"Qty" validate:"required"`
	// ClosePositionOrder は決済順序。0:日付古→損益高、1:日付古→損益低、2:日付新→損益高、3:日付新→損益低、
	// 4:損益高→日付古、5:損益高→日付新、6:損益低→日付古、7:損益低→日付新。ClosePositions との同時指定は不可。
	ClosePositionOrder int `json:"ClosePositionOrder,omitempty"`
	// ClosePositions は返済建玉を個別指定する。ClosePositionOrder と排他。
	ClosePositions []struct {
		HoldID string `json:"HoldID,omitempty"` // 返済建玉ID。
		Qty    int    `json:"Qty,omitempty"`    // 返済建玉数量。
	} `json:"ClosePositions,omitempty"`
	// FrontOrderType は執行条件。18:引成（TimeInForce は FAK のみ）、20:指値、28:引指（TimeInForce は FAS のみ）、
	// 30:逆指値（AfterHitPrice 側で価格を指定）、120:成行（Price は 0 を指定）。
	FrontOrderType int `json:"FrontOrderType" validate:"required"`
	// Price は注文価格。成行や引成では 0 を指定し、その他は発注したい単価を設定する。
	Price float64 `json:"Price" validate:"required"`
	// ExpireDay は注文有効期限の日付。yyyyMMdd 形式、0 は「本日」扱い（市場の引け後は翌取引所営業日、休前日は休日明け）。
	// 日替わりは kabu ステーション側の日付更新タイミングに従う。
	ExpireDay int `json:"ExpireDay" validate:"required"`
	// ReverseLimitOrder は FrontOrderType=逆指値のときに設定する条件。
	ReverseLimitOrder *struct {
		TriggerPrice      float64 `json:"TriggerPrice,omitempty"`      // TriggerPrice は逆指値の発動価格。未設定や数値以外はエラー。
		UnderOver         int     `json:"UnderOver,omitempty"`         // UnderOver は発動条件。1:以下、2:以上。未設定や 1/2 以外はエラー。
		AfterHitOrderType int     `json:"AfterHitOrderType,omitempty"` // AfterHitOrderType はヒット後の執行条件。1:成行、2:指値。日通しでは 2 のみ有効、日中/夜間では 1 または 2 を指定。逆指値成行は TimeInForce=FAK、逆指値指値は TimeInForce=FAS を指定する。
		AfterHitPrice     float64 `json:"AfterHitPrice"`               // AfterHitPrice はヒット後の価格。成行なら 0、指値なら単価を指定。未設定や数値以外はエラー。
	} `json:"ReverseLimitOrder,omitempty"`
}

// ResPostOrderSendorderFuture は **POST /sendorder/future** のレスポンス。
//
// ### 概要
// 注文発注（先物）
type ResPostOrderSendorderFuture struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PostOrderSendorderFuture は **POST /sendorder/future** を実行し、先物の注文を発注します。
//
// 機能:
//   - 取引区分や数量、執行条件などを指定し、先物取引の新規・返済注文を送信する。
//
// 引数:
//   - req (ReqPostOrderSendorderFuture): 先物注文の内容をまとめたリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPostOrderSendorderFuture): 受付結果と注文番号。
//   - err (error): 通信エラーまたは APIError。
func PostOrderSendorderFuture(req ReqPostOrderSendorderFuture) (code int, res ResPostOrderSendorderFuture, err error) {
	p := "/sendorder/future"
	v := url.Values{}
	// リクエストボディをJSON化
	b, err := json.Marshal(req)
	if err != nil {
		return 0, res, err
	}
	needAuth := true
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

// ReqPutOrderCancelorder は **PUT /cancelorder** のリクエスト。
//
// ### 概要
// 注文取消
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPutOrderCancelorder struct {
	// 注文番号<br>sendorderのレスポンスで受け取るOrderID。
	OrderId string `json:"OrderId" validate:"required"`
}

// ResPutOrderCancelorder は **PUT /cancelorder** のレスポンス。
//
// ### 概要
// 注文取消
type ResPutOrderCancelorder struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PutOrderCancelorder は **PUT /cancelorder** を実行し、既存注文を取り消します。
//
// 機能:
//   - 受付済みの注文番号を指定して、発注済みの注文をキャンセルする。
//
// 引数:
//   - req (ReqPutOrderCancelorder): 取消対象の OrderId を含むリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPutOrderCancelorder): 取消結果と注文番号。
//   - err (error): 通信エラーまたは APIError。
func PutOrderCancelorder(req ReqPutOrderCancelorder) (code int, res ResPutOrderCancelorder, err error) {
	p := "/cancelorder"
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
