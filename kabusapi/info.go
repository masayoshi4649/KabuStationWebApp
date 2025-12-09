package kabusapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ReqGetInfoOrders は **GET /orders** のリクエスト。
//
// 注文約定照会
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoOrders struct {
	// 取得する商品 定義値 説明 0 すべて 1 現物 2 信用 3 先物 4 OP
	Product string `json:"-" query:"product"`
	// 注文番号 ※指定された注文番号と一致する注文のみレスポンスします。 ※指定された注文番号との比較では大文字小文字を区別しません。 ※複数の注文番号を指定することはできません。
	ID string `json:"-" query:"id"`
	// 更新日時 ※形式：yyyyMMddHHmmss （例：20201201123456） ※指定された更新日時以降（指定日時含む）に更新された注文のみレスポンスします。 ※複数の更新日時を指定することはできません。
	Updtime string `json:"-" query:"updtime"`
	// 注文詳細抑止 定義値 説明 true 注文詳細を出力する（デフォルト） false 注文詳細の出力しない
	Details string `json:"-" query:"details"`
	// 銘柄コード ※指定された銘柄コードと一致する注文のみレスポンスします。 ※複数の銘柄コードを指定することができません。
	Symbol string `json:"-" query:"symbol"`
	// 状態 ※指定された状態と一致する注文のみレスポンスします。 ※フィルタには数字の入力のみ受け付けます。 ※複数の状態を指定することはできません。 定義値 説明 1 待機（発注待機） 2 処理中（発注送信中） 3 処理済（発注済・訂正済） 4 訂正取消送信中 5 終了（発注エラー・取消済・全約定・失効・期限切れ）
	State string `json:"-" query:"state"`
	// 売買区分 ※指定された売買区分と一致する注文のみレスポンスします。 ※フィルタには数字の入力のみ受け付けます。 ※複数の売買区分を指定することができません。 定義値 説明 1 売 2 買
	Side string `json:"-" query:"side"`
	// 取引区分 ※指定された取引区分と一致する注文のみレスポンスします。 ※フィルタには数字の入力のみ受け付けます。 ※複数の取引区分を指定することができません。 定義値 説明 2 新規 3 返済
	Cashmargin string `json:"-" query:"cashmargin"`
}

// ResGetInfoOrders は **GET /orders** のレスポンス。
//
// 注文約定照会
type ResGetInfoOrders []struct {
	// 注文番号
	ID string `json:"ID,omitempty"`
	// 状態 ※OrderState と同一。1:待機 2:処理中 3:処理済 4:訂正取消送信中 5:終了
	State int `json:"State,omitempty"`
	// 注文状態 ※State と同一。1:待機 2:処理中 3:処理済 4:訂正取消送信中 5:終了
	OrderState int `json:"OrderState,omitempty"`
	// 執行条件 1:ザラバ 2:寄り 3:引け 4:不成 5:対当指値 6:IOC
	OrdType int `json:"OrdType,omitempty"`
	// 受注日時
	RecvTime string `json:"RecvTime,omitempty"`
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名
	SymbolName string `json:"SymbolName,omitempty"`
	// 市場コード 1:東証 3:名証 5:福証 6:札証 9:SOR 27:東証+ 2:日通し 23:日中 24:夜間
	Exchange int `json:"Exchange,omitempty"`
	// 市場名
	ExchangeName string `json:"ExchangeName,omitempty"`
	// 有効期間条件 ※先物・オプション銘柄の場合のみ 1:FAS 2:FAK 3:FOK
	TimeInForce int `json:"TimeInForce,omitempty"`
	// 値段
	Price float64 `json:"Price,omitempty"`
	// 発注数量 ※期限切れや失効でもゼロにはならない
	OrderQty float64 `json:"OrderQty,omitempty"`
	// 約定数量
	CumQty float64 `json:"CumQty,omitempty"`
	// 売買区分 1:売 2:買
	Side string `json:"Side,omitempty"`
	// 取引区分 2:新規 3:返済
	CashMargin int `json:"CashMargin,omitempty"`
	// 口座種別 2:一般 4:特定 12:法人
	AccountType int `json:"AccountType,omitempty"`
	// 受渡区分 2:お預り金 3:auマネーコネクト
	DelivType int `json:"DelivType,omitempty"`
	// 注文有効期限 yyyyMMdd 形式
	ExpireDay int `json:"ExpireDay,omitempty"`
	// 信用取引区分 1:制度信用 2:一般信用（長期） 3:一般信用（デイトレ）
	MarginTradeType int `json:"MarginTradeType,omitempty"`
	// プレミアム料 ※信用を注文した際に表示
	MarginPremium float64 `json:"MarginPremium,omitempty"`
	// 注文詳細
	Details []struct {
		// 注文明細レコードの生成順序。通番ではないが大小による順序は維持される。
		SeqNum int `json:"SeqNum,omitempty"`
		// 注文詳細番号
		ID string `json:"ID,omitempty"`
		// 明細種別 1:受付 2:繰越 3:期限切れ 4:発注 5:訂正 6:取消 7:失効 8:約定
		RecType int `json:"RecType,omitempty"`
		// 取引所番号
		ExchangeID string `json:"ExchangeID,omitempty"`
		// 状態 1:待機 2:処理中 3:処理済 4:エラー 5:削除済み
		State int `json:"State,omitempty"`
		// 処理時刻
		TransactTime string `json:"TransactTime,omitempty"`
		// 執行条件 null:取消 / 0:期限切れ・失効・約定 / 1:ザラバ / 2:寄り / 3:引け / 4:不成 / 5:対当指値 / 6:IOC
		OrdType int `json:"OrdType,omitempty"`
		// 値段
		Price float64 `json:"Price,omitempty"`
		// 数量
		Qty float64 `json:"Qty,omitempty"`
		// 約定番号
		ExecutionID string `json:"ExecutionID,omitempty"`
		// 約定日時
		ExecutionDay string `json:"ExecutionDay,omitempty"`
		// 受渡日
		DelivDay int `json:"DelivDay,omitempty"`
		// 手数料 ※明細種別が約定（RecType=8）の場合に設定
		Commission float64 `json:"Commission,omitempty"`
		// 手数料消費税 ※明細種別が約定（RecType=8）の場合に設定
		CommissionTax float64 `json:"CommissionTax,omitempty"`
	} `json:"Details,omitempty"`
}

// GetInfoOrders は **GET /orders** を呼び出して注文・約定状況を取得します。
//
// 機能:
//   - 商品区分や注文番号などで絞り込み、現在の注文一覧を取得する。
//
// 引数:
//   - req (ReqGetInfoOrders): 注文照会の絞り込み条件。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoOrders): 注文・約定に関する情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoOrders(req ReqGetInfoOrders) (code int, res ResGetInfoOrders, err error) {
	p := "/orders"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.Product != "" {
		v.Set("product", fmt.Sprint(req.Product))
	}
	if req.ID != "" {
		v.Set("id", fmt.Sprint(req.ID))
	}
	if req.Updtime != "" {
		v.Set("updtime", fmt.Sprint(req.Updtime))
	}
	if req.Details != "" {
		v.Set("details", fmt.Sprint(req.Details))
	}
	if req.Symbol != "" {
		v.Set("symbol", fmt.Sprint(req.Symbol))
	}
	if req.State != "" {
		v.Set("state", fmt.Sprint(req.State))
	}
	if req.Side != "" {
		v.Set("side", fmt.Sprint(req.Side))
	}
	if req.Cashmargin != "" {
		v.Set("cashmargin", fmt.Sprint(req.Cashmargin))
	}
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

// ReqGetInfoPositions は **GET /positions** のリクエスト。
//
// 残高照会
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoPositions struct {
	// 取得する商品 定義値 説明 0 すべて 1 現物 2 信用 3 先物 4 OP
	Product string `json:"-" query:"product"`
	// 銘柄コード ※指定された銘柄コードと一致するポジションのみレスポンスします。 ※複数の銘柄コードを指定することはできません。
	Symbol string `json:"-" query:"symbol"`
	// 売買区分フィルタ 指定された売買区分と一致する注文を返す 定義値 説明 1 売 2 買
	Side string `json:"-" query:"side"`
	// 追加情報出力フラグ（未指定時：true） ※追加情報は、「現在値」、「評価金額」、「評価損益額」、「評価損益率」を意味します。 定義値 説明 true 追加情報を出力する false 追加情報を出力しない
	Addinfo string `json:"-" query:"addinfo"`
}

// ResGetInfoPositions は **GET /positions** のレスポンス。
//
// 残高照会
type ResGetInfoPositions []struct {
	// 約定番号 ※現物取引では、nullが返ります。
	ExecutionID string `json:"ExecutionID,omitempty"`
	// 口座種別 2:一般 4:特定 12:法人
	AccountType int `json:"AccountType,omitempty"`
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名
	SymbolName string `json:"SymbolName,omitempty"`
	// 市場コード 1:東証 3:名証 5:福証 6:札証 9:SOR 27:東証+ 2:日通し 23:日中 24:夜間
	Exchange int `json:"Exchange,omitempty"`
	// 市場名
	ExchangeName string `json:"ExchangeName,omitempty"`
	// 銘柄種別 ※先物・オプション銘柄の場合のみ 0:指数 1:現物 101:日経225先物 103:日経225OP 107:TOPIX先物 121:JPX400先物 144:NYダウ 145:日経平均VI 154:グロース250先物 155:TOPIX_REIT 171:TOPIX CORE30 901:日経平均225ミニ先物 907:TOPIXミニ先物
	SecurityType int `json:"SecurityType,omitempty"`
	// 約定日（建玉日） ※信用・先物・オプションの場合のみ ※現物取引では、nullが返ります。
	ExecutionDay int `json:"ExecutionDay,omitempty"`
	// 値段
	Price float64 `json:"Price,omitempty"`
	// 残数量（保有数量）
	LeavesQty float64 `json:"LeavesQty,omitempty"`
	// 拘束数量（返済のために拘束されている数量）
	HoldQty float64 `json:"HoldQty,omitempty"`
	// 売買区分 1:売 2:買
	Side string `json:"Side,omitempty"`
	// 諸経費 ※信用・先物・オプションの場合のみ
	Expenses float64 `json:"Expenses,omitempty"`
	// 手数料 ※信用・先物・オプションの場合のみ
	Commission float64 `json:"Commission,omitempty"`
	// 手数料消費税 ※信用・先物・オプションの場合のみ
	CommissionTax float64 `json:"CommissionTax,omitempty"`
	// 返済期日 ※信用・先物・オプションの場合のみ
	ExpireDay int `json:"ExpireDay,omitempty"`
	// 信用取引区分 ※信用の場合のみ 1:制度信用 2:一般信用（長期） 3:一般信用（デイトレ）
	MarginTradeType int `json:"MarginTradeType,omitempty"`
	// 現在値 追加情報出力フラグ：falseの場合、null
	CurrentPrice float64 `json:"CurrentPrice,omitempty"`
	// 評価金額 追加情報出力フラグ：falseの場合、null
	Valuation float64 `json:"Valuation,omitempty"`
	// 評価損益額 追加情報出力フラグ：falseの場合、null
	ProfitLoss float64 `json:"ProfitLoss,omitempty"`
	// 評価損益率 追加情報出力フラグ：falseの場合、null
	ProfitLossRate float64 `json:"ProfitLossRate,omitempty"`
}

// GetInfoPositions は **GET /positions** を呼び出して建玉残高を取得します。
//
// 機能:
//   - 商品区分や銘柄、保有区分でフィルタし、保有中の建玉一覧を取得する。
//
// 引数:
//   - req (ReqGetInfoPositions): 残高照会の絞り込み条件。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoPositions): 建玉残高の一覧。
//   - err (error): 通信エラーまたは APIError。
func GetInfoPositions(req ReqGetInfoPositions) (code int, res ResGetInfoPositions, err error) {
	p := "/positions"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.Product != "" {
		v.Set("product", fmt.Sprint(req.Product))
	}
	if req.Symbol != "" {
		v.Set("symbol", fmt.Sprint(req.Symbol))
	}
	if req.Side != "" {
		v.Set("side", fmt.Sprint(req.Side))
	}
	if req.Addinfo != "" {
		v.Set("addinfo", fmt.Sprint(req.Addinfo))
	}
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

// ReqGetInfoSymbolnameFuture は **GET /symbolname/future** のリクエスト。
//
// 先物銘柄コード取得
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolnameFuture struct {
	// 先物コード ※大文字小文字は区別しません。 定義値 説明 NK225 日経平均先物 NK225mini 日経225mini先物 TOPIX TOPIX先物 TOPIXmini ミニTOPIX先物 GROWTH グロース250先物 JPX400 JPX日経400先物 DOW NYダウ先物 VI 日経平均VI先物 Core30 TOPIX Core30先物 REIT 東証REIT指数先物 NK225micro 日経225マイクロ先物
	FutureCode string `json:"-" query:"FutureCode"`
	// 限月 ※限月はyyyyMM形式で指定します。0を指定した場合、直近限月となります。 ※取引最終日に「0」（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、 取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	DerivMonth string `json:"-" query:"DerivMonth" required:"true"`
}

// ResGetInfoSymbolnameFuture は **GET /symbolname/future** のレスポンス。
//
// 先物銘柄コード取得
type ResGetInfoSymbolnameFuture struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名称
	SymbolName string `json:"SymbolName,omitempty"`
}

// GetInfoSymbolnameFuture は **GET /symbolname/future** を呼び出して先物銘柄コードを取得します。
//
// 機能:
//   - 先物銘柄コード区分や限月指定に基づき、取引可能な先物銘柄コード一覧を取得する。
//
// 引数:
//   - req (ReqGetInfoSymbolnameFuture): 先物コード区分や限月、取引日を指定するリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoSymbolnameFuture): 取得した先物銘柄コード情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoSymbolnameFuture(req ReqGetInfoSymbolnameFuture) (code int, res ResGetInfoSymbolnameFuture, err error) {
	p := "/symbolname/future"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	v.Set("FutureCode", fmt.Sprint(req.FutureCode))
	v.Set("DerivMonth", fmt.Sprint(req.DerivMonth))

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
