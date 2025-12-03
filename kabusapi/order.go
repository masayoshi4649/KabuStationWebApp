package kabusapi

import (
	"encoding/json"
	"net/url"
)

// ReqPostOrderSendorder は **POST /sendorder** のリクエスト。
//
// ### 概要
// 注文発注（現物・信用）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPostOrderSendorder struct {
	// 銘柄コード
	Symbol string `json:"Symbol" validate:"required"`
	// 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>9</td> <td>SOR</td> </tr> <tr> <td>27</td> <td>東証+</td> </tr> </tbody> </table>
	Exchange int `json:"Exchange" validate:"required"`
	// 商品種別 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>株式</td> </tr> </tbody> </table>
	SecurityType int `json:"SecurityType" validate:"required"`
	// 売買区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
	Side string `json:"Side" validate:"required"`
	// 信用区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>現物</td> </tr> <tr> <td>2</td> <td>新規</td> </tr> <tr> <td>3</td> <td>返済</td> </tr> </tbody> </table>
	CashMargin int `json:"CashMargin" validate:"required"`
	// 信用取引区分<br>※現物取引の場合は指定不要。<br>※信用取引の場合、必須。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>制度信用</td> </tr> <tr> <td>2</td> <td>一般信用（長期）</td> </tr> <tr> <td>3</td> <td>一般信用（デイトレ）</td> </tr> </tbody> </table>
	MarginTradeType int `json:"MarginTradeType"`
	// １株あたりのプレミアム料(円)<br> ※プレミアム料の刻値は、プレミアム料取得APIのレスポンスにある"TickMarginPremium"にてご確認ください。<br> ※入札受付中(19:30～20:30)プレミアム料入札可能銘柄の場合、「MarginPremiumUnit」は必須となります。<br> ※入札受付中(19:30～20:30)のプレミアム料入札可能銘柄以外の場合は、「MarginPremiumUnit」の記載は無視されます。<br> ※入札受付中以外の時間帯では、「MarginPremiumUnit」の記載は無視されます。
	MarginPremiumUnit float64 `json:"MarginPremiumUnit"`
	// 受渡区分<br>※現物買は指定必須。<br>※現物売は「0(指定なし)」を設定<br>※信用新規は「0(指定なし)」を設定<br>※信用返済は指定必須 <br>※auマネーコネクトが有効の場合にのみ、「3」を設定可能 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>指定なし</td> </tr> <tr> <td>2</td> <td>お預り金</td> </tr> <tr> <td>3</td> <td>auマネーコネクト</td> </tr> </tbody> </table>
	DelivType int `json:"DelivType" validate:"required"`
	// 資産区分（預り区分）<br>※現物買は、指定必須。<br>※現物売は、「' '」 半角スペース2つを指定必須。<br>※信用新規と信用返済は、「11」を指定するか、または指定なしでも可。指定しない場合は「11」が自動的にセットされます。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>(半角スペース2つ)</td> <td>現物売の場合</td> </tr> <tr> <td>02</td> <td>保護</td> </tr> <tr> <td>AA</td> <td>信用代用</td> </tr> <tr> <td>11</td> <td>信用取引</td> </tr> </tbody> </table>
	FundType string `json:"FundType"`
	// 口座種別 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>一般</td> </tr> <tr> <td>4</td> <td>特定</td> </tr> <tr> <td>12</td> <td>法人</td> </tr> </tbody> </table>
	AccountType int `json:"AccountType" validate:"required"`
	// 注文数量<br>※信用一括返済の場合、返済したい合計数量を入力してください。
	Qty int `json:"Qty" validate:"required"`
	// 決済順序<br>※信用返済の場合、必須。<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>日付（古い順）、損益（高い順）</td> </tr> <tr> <td>1</td> <td>日付（古い順）、損益（低い順）</td> </tr> <tr> <td>2</td> <td>日付（新しい順）、損益（高い順）</td> </tr> <tr> <td>3</td> <td>日付（新しい順）、損益（低い順）</td> </tr> <tr> <td>4</td> <td>損益（高い順）、日付（古い順）</td> </tr> <tr> <td>5</td> <td>損益（高い順）、日付（新しい順）</td> </tr> <tr> <td>6</td> <td>損益（低い順）、日付（古い順）</td> </tr> <tr> <td>7</td> <td>損益（低い順）、日付（新しい順）</td> </tr> </tbody> </table>
	ClosePositionOrder int `json:"ClosePositionOrder"`
	// 返済建玉指定<br>※信用返済の場合、必須。<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。<br>※信用一括返済の場合、各建玉IDと返済したい数量を入力してください。<br>※建玉IDは「E」から始まる番号です。
	ClosePositions []struct {
		HoldID string `json:"HoldID,omitempty"` // 返済建玉ID
		Qty    int    `json:"Qty,omitempty"`    // 返済建玉数量
	} `json:"ClosePositions"`
	// 執行条件 ※SOR以外は以下、全て指定可能です。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> <th>”Price"の指定</th> <th>SORで発注可</th> </tr> </thead> <tbody> <tr> <td>10</td> <td>成行</td> <td>0</td> <td>〇</td> </tr> <tr> <td>13</td> <td>寄成（前場）</td> <td>0</td> <td> </td> </tr> <tr> <td>14</td> <td>寄成（後場）</td> <td>0</td> <td> </td> </tr> <tr> <td>15</td> <td>引成（前場）</td> <td>0</td> <td> </td> </tr> <tr> <td>16</td> <td>引成（後場）</td> <td>0</td> <td> </td> </tr> <tr> <td>17</td> <td>IOC成行</td> <td>0</td> <td> </td> </tr> <tr> <td>20</td> <td>指値</td> <td>発注したい金額</td> <td>〇</td> </tr> <tr> <td>21</td> <td>寄指（前場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>22</td> <td>寄指（後場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>23</td> <td>引指（前場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>24</td> <td>引指（後場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>25</td> <td>不成（前場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>26</td> <td>不成（後場）</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>27</td> <td>IOC指値</td> <td>発注したい金額</td> <td> </td> </tr> <tr> <td>30</td> <td>逆指値</td> <td>指定なし<br>※AfterHitPriceで指定ください</td> <td>〇</td> </tr> </tbody> </table>
	FrontOrderType int `json:"FrontOrderType" validate:"required"`
	// 注文価格<br>※FrontOrderTypeで成行を指定した場合、0を指定する。<br>※詳細について、”FrontOrderType”をご確認ください。
	Price float64 `json:"Price" validate:"required"`
	// 注文有効期限<br> yyyyMMdd形式。<br> 「0」を指定すると、kabuステーション上の発注画面の「本日」に対応する日付として扱います。<br> 「本日」は直近の注文可能日となり、以下のように設定されます。<br> 引けまでの間 : 当日<br> 引け後 : 翌取引所営業日<br> 休前日 : 休日明けの取引所営業日<br> ※ 日替わりはkabuステーションが日付変更通知を受信したタイミングです。
	ExpireDay int `json:"ExpireDay" validate:"required"`
	// 逆指値条件<br> ※FrontOrderTypeで逆指値を指定した場合のみ必須。
	ReverseLimitOrder struct {
		TriggerSec        int     `json:"TriggerSec,omitempty"`        // トリガ銘柄<br> ※未設定の場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>発注銘柄</td> </tr> <tr> <td>2</td> <td>NK225指数</td> </tr> <tr> <td>3</td> <td>TOPIX指数</td> </tr> </tbody> </table>
		TriggerPrice      float64 `json:"TriggerPrice,omitempty"`      // トリガ価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。
		UnderOver         int     `json:"UnderOver,omitempty"`         // 以上／以下<br> ※未設定の場合はエラーになります。<br> ※1、2以外が指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>以下</td> </tr> <tr> <td>2</td> <td>以上</td> </tr> </tbody> </table>
		AfterHitOrderType int     `json:"AfterHitOrderType,omitempty"` // ヒット後執行条件<br> ※未設定の場合はエラーになります。<br> ※1、2、3以外が指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>成行</td> </tr> <tr> <td>2</td> <td>指値</td> </tr> <tr> <td>3</td> <td>不成</td> </tr> </tbody> </table>
		AfterHitPrice     float64 `json:"AfterHitPrice,omitempty"`     // ヒット後注文価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。<br><br> ヒット後執行条件に従い、下記のようにヒット後注文価格を設定してください。 <table> <thead> <tr> <th>ヒット後執行条件</th> <th>設定価格</th> </tr> </thead> <tbody> <tr> <td>成行</td> <td>0</td> </tr> <tr> <td>指値</td> <td>指値の単価</td> </tr> <tr> <td>不成</td> <td>不成の単価</td> </tr> </tbody> </table>
	} `json:"ReverseLimitOrder"`
}

// ResPostOrderSendorder は **POST /sendorder** のレスポンス。
//
// ### 概要
// 注文発注（現物・信用）
type ResPostOrderSendorder struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PostOrderSendorder は **POST /sendorder** を実行し、現物・信用の注文を発注します。
//
// 機能:
//   - 銘柄コード、売買区分、数量、価格、執行条件などを指定して株式（現物/信用）の新規・返済注文を送信する。
//
// 引数:
//   - req (ReqPostOrderSendorder): 発注内容をまとめたリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPostOrderSendorder): 受付結果と注文番号。
//   - err (error): 通信エラーまたは APIError。
func PostOrderSendorder(req ReqPostOrderSendorder) (code int, res ResPostOrderSendorder, err error) {
	p := "/sendorder"
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
	ReverseLimitOrder struct {
		TriggerPrice      float64 `json:"TriggerPrice,omitempty"`      // TriggerPrice は逆指値の発動価格。未設定や数値以外はエラー。
		UnderOver         int     `json:"UnderOver,omitempty"`         // UnderOver は発動条件。1:以下、2:以上。未設定や 1/2 以外はエラー。
		AfterHitOrderType int     `json:"AfterHitOrderType,omitempty"` // AfterHitOrderType はヒット後の執行条件。1:成行、2:指値。日通しでは 2 のみ有効、日中/夜間では 1 または 2 を指定。逆指値成行は TimeInForce=FAK、逆指値指値は TimeInForce=FAS を指定する。
		AfterHitPrice     float64 `json:"AfterHitPrice,omitempty"`     // AfterHitPrice はヒット後の価格。成行なら 0、指値なら単価を指定。未設定や数値以外はエラー。
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

// ReqPostOrderSendorderOption は **POST /sendorder/option** のリクエスト。
//
// ### 概要
// 注文発注（オプション）
//
// - 認証: 必須（X-API-KEY ヘッダー）
type ReqPostOrderSendorderOption struct {
	// 銘柄コード<br>※取引最終日に「オプション銘柄コード取得」でDerivMonthに0（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	Symbol string `json:"Symbol" validate:"required"`
	// 市場コード <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>2</td> <td>日通し</td> </tr> <tr> <td>23</td> <td>日中</td> </tr> <tr> <td>24</td> <td>夜間</td> </tr> </tbody> </table>
	Exchange int `json:"Exchange" validate:"required"`
	// 取引区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>新規</td> </tr> <tr> <td>2</td> <td>返済</td> </tr> </tbody> </table>
	TradeType int `json:"TradeType" validate:"required"`
	// 有効期間条件 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>FAS</td> </tr> <tr> <td>2</td> <td>FAK</td> </tr> <tr> <td>3</td> <td>FOK</td> </tr> </tbody> </table> ※執行条件(FrontOrderType)、有効期限条件(TimeInForce)、市場コード(Exchange)で選択できる組み合わせは下表のようになります。 <table> <thead> <tr> <th rowspan="2">執行条件</th> <th rowspan="2">有効期間条件</th> <th colspan="3">市場コード</th> </tr> <tr> <th>日中</th> <th>夜間</th> <th>日通し</th> </tr> </thead> <tbody> <tr> <td>指値</td> <td>FAS</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>指値</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>指値</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>成行</td> <td>FOK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>逆指値（指値）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>●</td> </tr> <tr> <td>逆指値（成行）</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>引成</td> <td>FAK</td> <td>●</td> <td>●</td> <td>-</td> </tr> <tr> <td>引指</td> <td>FAS</td> <td>●</td> <td>●</td> <td>-</td> </tr> </tbody> </table>
	TimeInForce int `json:"TimeInForce" validate:"required"`
	// 売買区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
	Side string `json:"Side" validate:"required"`
	// 注文数量
	Qty int `json:"Qty" validate:"required"`
	// 決済順序<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>日付（古い順）、損益（高い順）</td> </tr> <tr> <td>1</td> <td>日付（古い順）、損益（低い順）</td> </tr> <tr> <td>2</td> <td>日付（新しい順）、損益（高い順）</td> </tr> <tr> <td>3</td> <td>日付（新しい順）、損益（低い順）</td> </tr> <tr> <td>4</td> <td>損益（高い順）、日付（古い順）</td> </tr> <tr> <td>5</td> <td>損益（高い順）、日付（新しい順）</td> </tr> <tr> <td>6</td> <td>損益（低い順）、日付（古い順）</td> </tr> <tr> <td>7</td> <td>損益（低い順）、日付（新しい順）</td> </tr> </tbody> </table>
	ClosePositionOrder int `json:"ClosePositionOrder"`
	// 返済建玉指定<br>※ClosePositionOrderとClosePositionsはどちらか一方のみ指定可能。<br>※ClosePositionOrderとClosePositionsを両方指定した場合、エラー。
	ClosePositions []struct {
		HoldID string `json:"HoldID,omitempty"` // 返済建玉ID
		Qty    int    `json:"Qty,omitempty"`    // 返済建玉数量
	} `json:"ClosePositions"`
	// 執行条件 <table> <thead> <tr> <th>定義値</th> <th>説明</th> <th>”Price”の指定</th> </tr> </thead> <tbody> <tr> <td>18</td> <td>引成（派生）<br>※TimeInForceは、「FAK」のみ有効</td> <td>0</td> </tr> <tr> <td>20</td> <td>指値</td> <td>発注したい金額</td> </tr> <tr> <td>28</td> <td>引指（派生）<br>※TimeInForceは、「FAS」のみ有効</td> <td>発注したい金額</td> </tr> <tr> <td>30</td> <td>逆指値</td> <td>指定なし<br>※AfterHitPriceで指定ください</td> </tr> <tr> <td>120</td> <td>成行（マーケットオーダー）</td> <td>0</td> </tr> </tbody> </table>
	FrontOrderType int `json:"FrontOrderType" validate:"required"`
	// 注文価格<br>※FrontOrderTypeで成行を指定した場合、0を指定する。<br>※詳細について、”FrontOrderType”をご確認ください。
	Price float64 `json:"Price" validate:"required"`
	// 注文有効期限<br> yyyyMMdd形式。<br> 「0」を指定すると、kabuステーション上の発注画面の「本日」に対応する日付として扱います。<br> 「本日」は直近の注文可能日となり、以下のように設定されます。<br> その市場の引けまでの間 : 当日<br> その市場の引け後 : 翌取引所営業日<br> その市場の休前日 : 休日明けの取引所営業日<br> ※ 日替わりはkabuステーションが日付変更通知を受信したタイミングです。<br> ※ 日通しの場合、夜間取引の引け後に日付が更新されます。
	ExpireDay int `json:"ExpireDay" validate:"required"`
	// 逆指値条件<br> ※FrontOrderTypeで逆指値を指定した場合のみ必須。
	ReverseLimitOrder struct {
		TriggerPrice      float64 `json:"TriggerPrice,omitempty"`      // トリガ価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。
		UnderOver         int     `json:"UnderOver,omitempty"`         // 以上／以下<br> ※未設定の場合はエラーになります。<br> ※1、2以外が指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>以下</td> </tr> <tr> <td>2</td> <td>以上</td> </tr> </tbody> </table>
		AfterHitOrderType int     `json:"AfterHitOrderType,omitempty"` // ヒット後執行条件<br> ※未設定の場合はエラーになります。<br> ※日通の注文で2以外が指定された場合はエラーになります。<br> ※日中、夜間の注文で1、2以外が指定された場合はエラーになります。<br> ※逆指値（成行）で有効期間条件(TimeInForce)にFAK以外を指定された場合はエラーになります。<br> ※逆指値（指値）で有効期間条件(TimeInForce)にFAS以外を指定された場合はエラーになります。 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>1</td> <td>成行</td> </tr> <tr> <td>2</td> <td>指値</td> </tr> </tbody> </table>
		AfterHitPrice     float64 `json:"AfterHitPrice,omitempty"`     // ヒット後注文価格<br> ※未設定の場合はエラーになります。<br> ※数字以外が設定された場合はエラーになります。<br><br> ヒット後執行条件に従い、下記のようにヒット後注文価格を設定してください。 <table> <thead> <tr> <th>ヒット後執行条件</th> <th>設定価格</th> </tr> </thead> <tbody> <tr> <td>成行</td> <td>0</td> </tr> <tr> <td>指値</td> <td>指値の単価</td> </tr> </tbody> </table>
	} `json:"ReverseLimitOrder"`
}

// ResPostOrderSendorderOption は **POST /sendorder/option** のレスポンス。
//
// ### 概要
// 注文発注（オプション）
type ResPostOrderSendorderOption struct {
	// 結果コード<br>0が成功。それ以外はエラーコード。
	Result int `json:"Result,omitempty"`
	// 受付注文番号
	OrderId string `json:"OrderId,omitempty"`
}

// PostOrderSendorderOption は **POST /sendorder/option** を実行し、オプションの注文を発注します。
//
// 機能:
//   - 売買区分や建玉指定、価格、執行条件などを指定し、オプション取引の新規・返済注文を送信する。
//
// 引数:
//   - req (ReqPostOrderSendorderOption): オプション注文の内容をまとめたリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResPostOrderSendorderOption): 受付結果と注文番号。
//   - err (error): 通信エラーまたは APIError。
func PostOrderSendorderOption(req ReqPostOrderSendorderOption) (code int, res ResPostOrderSendorderOption, err error) {
	p := "/sendorder/option"
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
