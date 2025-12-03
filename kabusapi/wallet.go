package kabusapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ReqGetWalletWalletCash は **GET /wallet/cash** のリクエスト。
//
// ### 概要
// 取引余力（現物）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletCash struct {
}

// ResGetWalletWalletCash は **GET /wallet/cash** のレスポンス。
//
// ### 概要
// 取引余力（現物）
type ResGetWalletWalletCash struct {
	// 現物買付可能額<br>※auマネーコネクトが有効の場合、auじぶん銀行の残高を含めた合計可能額を表示する<br>※auマネーコネクトが無効の場合、三菱UFJ eスマート証券の可能額のみを表示する
	StockAccountWallet float64 `json:"StockAccountWallet,omitempty"`
	// うち、三菱UFJ eスマート証券可能額
	AuKCStockAccountWallet float64 `json:"AuKCStockAccountWallet,omitempty"`
	// うち、auじぶん銀行残高<br>※auマネーコネクトが無効の場合、「0」を表示する
	AuJbnStockAccountWallet float64 `json:"AuJbnStockAccountWallet,omitempty"`
}

// GetWalletWalletCash は **GET /wallet/cash** を呼び出して現物の取引余力を取得します。
//
// 機能:
//   - 現物取引用の買付可能額を取得する。
//
// 引数:
//   - req (ReqGetWalletWalletCash): 空のリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletCash): 現物取引余力の情報。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletCash(req ReqGetWalletWalletCash) (code int, res ResGetWalletWalletCash, err error) {
	p := "/wallet/cash"
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetWalletWalletCashSymbol は **GET /wallet/cash/{symbol}** のリクエスト。
//
// ### 概要
// 取引余力（現物）（銘柄指定）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletCashSymbol struct {
	// 銘柄コード <br> ※次の形式で入力してください。<br> [銘柄コード]@[市場コード]<br> ※市場コードは下記の定義値から選択してください。 <b>市場コード</b> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>9</td> <td>SOR</td> </tr> <tr> <td>27</td> <td>東証+</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
}

// ResGetWalletWalletCashSymbol は **GET /wallet/cash/{symbol}** のレスポンス。
//
// ### 概要
// 取引余力（現物）（銘柄指定）
type ResGetWalletWalletCashSymbol struct {
	// 現物買付可能額<br>※auマネーコネクトが有効の場合、auじぶん銀行の残高を含めた合計可能額を表示する<br>※auマネーコネクトが無効の場合、三菱UFJ eスマート証券の可能額のみを表示する
	StockAccountWallet float64 `json:"StockAccountWallet,omitempty"`
	// うち、三菱UFJ eスマート証券可能額
	AuKCStockAccountWallet float64 `json:"AuKCStockAccountWallet,omitempty"`
	// うち、auじぶん銀行残高<br>※auマネーコネクトが無効の場合、「0」を表示する
	AuJbnStockAccountWallet float64 `json:"AuJbnStockAccountWallet,omitempty"`
}

// GetWalletWalletCashSymbol は **GET /wallet/cash/{symbol}** を呼び出して銘柄別の現物取引余力を取得します。
//
// 機能:
//   - 指定した銘柄コードに対する現物買付可能額を取得する。
//
// 引数:
//   - req (ReqGetWalletWalletCashSymbol): 銘柄コードと市場コードを含むリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletCashSymbol): 銘柄別の現物取引余力。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletCashSymbol(req ReqGetWalletWalletCashSymbol) (code int, res ResGetWalletWalletCashSymbol, err error) {
	p := "/wallet/cash/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetWalletWalletMargin は **GET /wallet/margin** のリクエスト。
//
// ### 概要
// 取引余力（信用）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletMargin struct {
}

// ResGetWalletWalletMargin は **GET /wallet/margin** のレスポンス。
//
// ### 概要
// 取引余力（信用）
type ResGetWalletWalletMargin struct {
	// 信用新規可能額
	MarginAccountWallet float64 `json:"MarginAccountWallet,omitempty"`
	// 保証金維持率<br>※銘柄指定の場合のみ<br>※銘柄が指定されなかった場合、0.0を返す。
	DepositkeepRate float64 `json:"DepositkeepRate,omitempty"`
	// 委託保証金率<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、Noneを返す。
	ConsignmentDepositRate float64 `json:"ConsignmentDepositRate,omitempty"`
	// 現金委託保証金率<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、Noneを返す。
	CashOfConsignmentDepositRate float64 `json:"CashOfConsignmentDepositRate,omitempty"`
}

// GetWalletWalletMargin は **GET /wallet/margin** を呼び出して信用取引余力を取得します。
//
// 機能:
//   - 信用新規可能額や保証金率など、信用取引に関する口座余力を取得する。
//
// 引数:
//   - req (ReqGetWalletWalletMargin): 空のリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletMargin): 信用取引余力の情報。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletMargin(req ReqGetWalletWalletMargin) (code int, res ResGetWalletWalletMargin, err error) {
	p := "/wallet/margin"
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetWalletWalletMarginSymbol は **GET /wallet/margin/{symbol}** のリクエスト。
//
// ### 概要
// 取引余力（信用）（銘柄指定）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletMarginSymbol struct {
	// 銘柄コード <br> ※次の形式で入力してください。<br> [銘柄コード]@[市場コード]<br> ※市場コードは下記の定義値から選択してください。 <b>市場コード</b> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>9</td> <td>SOR</td> </tr> <tr> <td>27</td> <td>東証+</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
}

// ResGetWalletWalletMarginSymbol は **GET /wallet/margin/{symbol}** のレスポンス。
//
// ### 概要
// 取引余力（信用）（銘柄指定）
type ResGetWalletWalletMarginSymbol struct {
	// 信用新規可能額
	MarginAccountWallet float64 `json:"MarginAccountWallet,omitempty"`
	// 保証金維持率<br>※銘柄指定の場合のみ<br>※銘柄が指定されなかった場合、0.0を返す。
	DepositkeepRate float64 `json:"DepositkeepRate,omitempty"`
	// 委託保証金率<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、Noneを返す。
	ConsignmentDepositRate float64 `json:"ConsignmentDepositRate,omitempty"`
	// 現金委託保証金率<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、Noneを返す。
	CashOfConsignmentDepositRate float64 `json:"CashOfConsignmentDepositRate,omitempty"`
}

// GetWalletWalletMarginSymbol は **GET /wallet/margin/{symbol}** を呼び出して銘柄別の信用取引余力を取得します。
//
// 機能:
//   - 指定した銘柄の保証金率や信用新規可能額などを取得する。
//
// 引数:
//   - req (ReqGetWalletWalletMarginSymbol): 銘柄コードと市場コードを含むリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletMarginSymbol): 銘柄別の信用取引余力。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletMarginSymbol(req ReqGetWalletWalletMarginSymbol) (code int, res ResGetWalletWalletMarginSymbol, err error) {
	p := "/wallet/margin/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetWalletWalletFuture は **GET /wallet/future** のリクエスト。
//
// ### 概要
// 取引余力（先物）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletFuture struct {
}

// ResGetWalletWalletFuture は **GET /wallet/future** のレスポンス。
//
// ### 概要
// 取引余力（先物）
type ResGetWalletWalletFuture struct {
	// 新規建玉可能額
	FutureTradeLimit float64 `json:"FutureTradeLimit,omitempty"`
	// 買い必要証拠金額<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、空を返す。
	MarginRequirement float64 `json:"MarginRequirement,omitempty"`
	// 売り必要証拠金額<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、空を返す。
	MarginRequirementSell float64 `json:"MarginRequirementSell,omitempty"`
}

// GetWalletWalletFuture は **GET /wallet/future** を呼び出して先物取引の余力を取得します。
//
// 機能:
//   - 先物口座の建玉評価額や拘束証拠金などを取得する。
//
// 引数:
//   - req (ReqGetWalletWalletFuture): 空のリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletFuture): 先物取引余力の情報。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletFuture(req ReqGetWalletWalletFuture) (code int, res ResGetWalletWalletFuture, err error) {
	p := "/wallet/future"
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetWalletWalletFutureSymbol は **GET /wallet/future/{symbol}** のリクエスト。
//
// ### 概要
// 取引余力（先物）（銘柄指定）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletFutureSymbol struct {
	// 銘柄コード <br> ※次の形式で入力してください。<br> [銘柄コード]@[市場コード]<br> ※市場コードは下記の定義値から選択してください。 ※SOR市場は取扱っておりませんのでご注意ください。<b>市場コード</b><br> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
}

// ResGetWalletWalletFutureSymbol は **GET /wallet/future/{symbol}** のレスポンス。
//
// ### 概要
// 取引余力（先物）（銘柄指定）
type ResGetWalletWalletFutureSymbol struct {
	// 新規建玉可能額
	FutureTradeLimit float64 `json:"FutureTradeLimit,omitempty"`
	// 買い必要証拠金額<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、空を返す。
	MarginRequirement float64 `json:"MarginRequirement,omitempty"`
	// 売り必要証拠金額<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、空を返す。
	MarginRequirementSell float64 `json:"MarginRequirementSell,omitempty"`
}

// GetWalletWalletFutureSymbol は **GET /wallet/future/{symbol}** を呼び出して銘柄別の先物取引余力を取得します。
//
// 機能:
//   - 指定した先物銘柄に対する必要証拠金や評価額を取得する。
//
// 引数:
//   - req (ReqGetWalletWalletFutureSymbol): 銘柄コードと市場コードを含むリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletFutureSymbol): 銘柄別の先物取引余力。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletFutureSymbol(req ReqGetWalletWalletFutureSymbol) (code int, res ResGetWalletWalletFutureSymbol, err error) {
	p := "/wallet/future/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetWalletWalletOption は **GET /wallet/option** のリクエスト。
//
// ### 概要
// 取引余力（オプション）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletOption struct {
}

// ResGetWalletWalletOption は **GET /wallet/option** のレスポンス。
//
// ### 概要
// 取引余力（オプション）
type ResGetWalletWalletOption struct {
	// 買新規建玉可能額
	OptionBuyTradeLimit float64 `json:"OptionBuyTradeLimit,omitempty"`
	// 売新規建玉可能額
	OptionSellTradeLimit float64 `json:"OptionSellTradeLimit,omitempty"`
	// 必要証拠金額<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、空を返す。
	MarginRequirement float64 `json:"MarginRequirement,omitempty"`
}

// GetWalletWalletOption は **GET /wallet/option** を呼び出してオプション取引余力を取得します。
//
// 機能:
//   - オプション取引における買建・売建可能額や必要証拠金を取得する。
//
// 引数:
//   - req (ReqGetWalletWalletOption): 空のリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletOption): オプション取引余力の情報。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletOption(req ReqGetWalletWalletOption) (code int, res ResGetWalletWalletOption, err error) {
	p := "/wallet/option"
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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

// ReqGetWalletWalletOptionSymbol は **GET /wallet/option/{symbol}** のリクエスト。
//
// ### 概要
// 取引余力（オプション）（銘柄指定）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetWalletWalletOptionSymbol struct {
	// 銘柄コード <br> ※次の形式で入力してください。<br> [銘柄コード]@[市場コード]<br> ※市場コードは下記の定義値から選択してください。 <b>市場コード</b> <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Symbol string `json:"-" path:"symbol"`
}

// ResGetWalletWalletOptionSymbol は **GET /wallet/option/{symbol}** のレスポンス。
//
// ### 概要
// 取引余力（オプション）（銘柄指定）
type ResGetWalletWalletOptionSymbol struct {
	// 買新規建玉可能額
	OptionBuyTradeLimit float64 `json:"OptionBuyTradeLimit,omitempty"`
	// 売新規建玉可能額
	OptionSellTradeLimit float64 `json:"OptionSellTradeLimit,omitempty"`
	// 必要証拠金額<br>※銘柄指定の場合のみ。<br>※銘柄が指定されなかった場合、空を返す。
	MarginRequirement float64 `json:"MarginRequirement,omitempty"`
}

// GetWalletWalletOptionSymbol は **GET /wallet/option/{symbol}** を呼び出して銘柄別のオプション取引余力を取得します。
//
// 機能:
//   - 指定したオプション銘柄の建玉可能額や必要証拠金を取得する。
//
// 引数:
//   - req (ReqGetWalletWalletOptionSymbol): 銘柄コードと市場コードを含むリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetWalletWalletOptionSymbol): 銘柄別のオプション取引余力。
//   - err (error): 通信エラーまたは APIError。
func GetWalletWalletOptionSymbol(req ReqGetWalletWalletOptionSymbol) (code int, res ResGetWalletWalletOptionSymbol, err error) {
	p := "/wallet/option/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	v := url.Values{}
	var b []byte // ボディなし
	needAuth := true
	code, data, err := doRequest("GET", p, v, b, needAuth)
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
