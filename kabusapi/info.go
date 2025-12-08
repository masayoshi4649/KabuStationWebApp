package kabusapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ReqGetInfoBoardSymbol は **GET /board/{symbol}** のリクエスト。
//
// 時価情報・板情報
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoBoardSymbol struct {
	// 銘柄コード ※次の形式で入力してください。 [銘柄コード]@[市場コード] ※市場コードは下記の定義値から選択してください。 ※SOR市場は取扱っておりませんのでご注意ください。市場コード 定義値 説明 1 東証 3 名証 5 福証 6 札証 2 日通し 23 日中 24 夜間
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoBoardSymbol は **GET /board/{symbol}** のレスポンス。
//
// 時価情報・板情報
type ResGetInfoBoardSymbol struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名
	SymbolName string `json:"SymbolName,omitempty"`
	// 市場コード ※株式・先物・オプション銘柄の場合のみ 定義値 説明 1 東証 3 名証 5 福証 6 札証 2 日通し 23 日中 24 夜間
	Exchange int `json:"Exchange,omitempty"`
	// 市場名称 ※株式・先物・オプション銘柄の場合のみ
	ExchangeName string `json:"ExchangeName,omitempty"`
	// 現値
	CurrentPrice float64 `json:"CurrentPrice,omitempty"`
	// 現値時刻
	CurrentPriceTime string `json:"CurrentPriceTime,omitempty"`
	// 現値前値比較 定義値 説明 0000 事象なし 0056 変わらず 0057 UP 0058 DOWN 0059 中断板寄り後の初値 0060 ザラバ引け 0061 板寄り引け 0062 中断引け 0063 ダウン引け 0064 逆転終値 0066 特別気配引け 0067 一時留保引け 0068 売買停止引け 0069 サーキットブレーカ引け 0431 ダイナミックサーキットブレーカ引け
	CurrentPriceChangeStatus string `json:"CurrentPriceChangeStatus,omitempty"`
	// 現値ステータス 定義値 説明 1 現値 2 不連続歩み 3 板寄せ 4 システム障害 5 中断 6 売買停止 7 売買停止・システム停止解除 8 終値 9 システム停止 10 概算値 11 参考値 12 サーキットブレイク実施中 13 システム障害解除 14 サーキットブレイク解除 15 中断解除 16 一時留保中 17 一時留保解除 18 ファイル障害 19 ファイル障害解除 20 Spread/Strategy 21 ダイナミックサーキットブレイク発動 22 ダイナミックサーキットブレイク解除 23 板寄せ約定
	CurrentPriceStatus int `json:"CurrentPriceStatus,omitempty"`
	// 計算用現値
	CalcPrice float64 `json:"CalcPrice,omitempty"`
	// 前日終値
	PreviousClose float64 `json:"PreviousClose,omitempty"`
	// 前日終値日付
	PreviousCloseTime string `json:"PreviousCloseTime,omitempty"`
	// 前日比
	ChangePreviousClose float64 `json:"ChangePreviousClose,omitempty"`
	// 騰落率
	ChangePreviousClosePer float64 `json:"ChangePreviousClosePer,omitempty"`
	// 始値
	OpeningPrice float64 `json:"OpeningPrice,omitempty"`
	// 始値時刻
	OpeningPriceTime string `json:"OpeningPriceTime,omitempty"`
	// 高値
	HighPrice float64 `json:"HighPrice,omitempty"`
	// 高値時刻
	HighPriceTime string `json:"HighPriceTime,omitempty"`
	// 安値
	LowPrice float64 `json:"LowPrice,omitempty"`
	// 安値時刻
	LowPriceTime string `json:"LowPriceTime,omitempty"`
	// 売買高 ※株式・先物・オプション銘柄の場合のみ
	TradingVolume float64 `json:"TradingVolume,omitempty"`
	// 売買高時刻 ※株式・先物・オプション銘柄の場合のみ
	TradingVolumeTime string `json:"TradingVolumeTime,omitempty"`
	// 売買高加重平均価格（VWAP） ※株式・先物・オプション銘柄の場合のみ
	VWAP float64 `json:"VWAP,omitempty"`
	// 売買代金 ※株式・先物・オプション銘柄の場合のみ
	TradingValue float64 `json:"TradingValue,omitempty"`
	// 最良売気配数量 ※① ※株式・先物・オプション銘柄の場合のみ
	BidQty float64 `json:"BidQty,omitempty"`
	// 最良売気配値段 ※① ※株式・先物・オプション銘柄の場合のみ
	BidPrice float64 `json:"BidPrice,omitempty"`
	// 最良売気配時刻 ※① ※株式銘柄の場合のみ
	BidTime string `json:"BidTime,omitempty"`
	// 最良売気配フラグ ※① ※株式・先物・オプション銘柄の場合のみ 定義値 説明 0000 事象なし 0101 一般気配 0102 特別気配 0103 注意気配 0107 寄前気配 0108 停止前特別気配 0109 引け後気配 0116 寄前気配約定成立ポイントなし 0117 寄前気配約定成立ポイントあり 0118 連続約定気配 0119 停止前の連続約定気配 0120 買い上がり売り下がり中
	BidSign string `json:"BidSign,omitempty"`
	// 売成行数量 ※株式銘柄の場合のみ
	MarketOrderSellQty float64 `json:"MarketOrderSellQty,omitempty"`
	// 売気配数量1本目
	Sell1 struct {
		Time  string  `json:"Time,omitempty"`  // 時刻<br>※株式銘柄の場合のみ
		Sign  string  `json:"Sign,omitempty"`  // 気配フラグ<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0000</td> <td>事象なし</td> </tr> <tr> <td>0101</td> <td>一般気配</td> </tr> <tr> <td>0102</td> <td>特別気配</td> </tr> <tr> <td>0103</td> <td>注意気配</td> </tr> <tr> <td>0107</td> <td>寄前気配</td> </tr> <tr> <td>0108</td> <td>停止前特別気配</td> </tr> <tr> <td>0109</td> <td>引け後気配</td> </tr> <tr> <td>0116</td> <td>寄前気配約定成立ポイントなし</td> </tr> <tr> <td>0117</td> <td>寄前気配約定成立ポイントあり</td> </tr> <tr> <td>0118</td> <td>連続約定気配</td> </tr> <tr> <td>0119</td> <td>停止前の連続約定気配</td> </tr> <tr> <td>0120</td> <td>買い上がり売り下がり中</td> </tr> </tbody> </table>
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell1,omitempty"`
	// 売気配数量2本目
	Sell2 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell2,omitempty"`
	// 売気配数量3本目
	Sell3 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell3,omitempty"`
	// 売気配数量4本目
	Sell4 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell4,omitempty"`
	// 売気配数量5本目
	Sell5 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell5,omitempty"`
	// 売気配数量6本目
	Sell6 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell6,omitempty"`
	// 売気配数量7本目
	Sell7 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell7,omitempty"`
	// 売気配数量8本目
	Sell8 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell8,omitempty"`
	// 売気配数量9本目
	Sell9 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell9,omitempty"`
	// 売気配数量10本目
	Sell10 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Sell10,omitempty"`
	// 最良買気配数量 ※① ※株式・先物・オプション銘柄の場合のみ
	AskQty float64 `json:"AskQty,omitempty"`
	// 最良買気配値段 ※① ※株式・先物・オプション銘柄の場合のみ
	AskPrice float64 `json:"AskPrice,omitempty"`
	// 最良買気配時刻 ※① ※株式銘柄の場合のみ
	AskTime string `json:"AskTime,omitempty"`
	// 最良買気配フラグ ※① ※株式・先物・オプション銘柄の場合のみ 定義値 説明 0000 事象なし 0101 一般気配 0102 特別気配 0103 注意気配 0107 寄前気配 0108 停止前特別気配 0109 引け後気配 0116 寄前気配約定成立ポイントなし 0117 寄前気配約定成立ポイントあり 0118 連続約定気配 0119 停止前の連続約定気配 0120 買い上がり売り下がり中
	AskSign string `json:"AskSign,omitempty"`
	// 買成行数量 ※株式銘柄の場合のみ
	MarketOrderBuyQty float64 `json:"MarketOrderBuyQty,omitempty"`
	// 買気配数量1本目
	Buy1 struct {
		Time  string  `json:"Time,omitempty"`  // 時刻<br>※株式銘柄の場合のみ
		Sign  string  `json:"Sign,omitempty"`  // 気配フラグ<br>※株式・先物・オプション銘柄の場合のみ <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>0000</td> <td>事象なし</td> </tr> <tr> <td>0101</td> <td>一般気配</td> </tr> <tr> <td>0102</td> <td>特別気配</td> </tr> <tr> <td>0103</td> <td>注意気配</td> </tr> <tr> <td>0107</td> <td>寄前気配</td> </tr> <tr> <td>0108</td> <td>停止前特別気配</td> </tr> <tr> <td>0109</td> <td>引け後気配</td> </tr> <tr> <td>0116</td> <td>寄前気配約定成立ポイントなし</td> </tr> <tr> <td>0117</td> <td>寄前気配約定成立ポイントあり</td> </tr> <tr> <td>0118</td> <td>連続約定気配</td> </tr> <tr> <td>0119</td> <td>停止前の連続約定気配</td> </tr> <tr> <td>0120</td> <td>買い上がり売り下がり中</td> </tr> </tbody> </table>
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy1,omitempty"`
	// 買気配数量2本目
	Buy2 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy2,omitempty"`
	// 買気配数量3本目
	Buy3 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy3,omitempty"`
	// 買気配数量4本目
	Buy4 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy4,omitempty"`
	// 買気配数量5本目
	Buy5 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy5,omitempty"`
	// 買気配数量6本目
	Buy6 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy6,omitempty"`
	// 買気配数量7本目
	Buy7 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy7,omitempty"`
	// 買気配数量8本目
	Buy8 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy8,omitempty"`
	// 買気配数量9本目
	Buy9 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy9,omitempty"`
	// 買気配数量10本目
	Buy10 struct {
		Price float64 `json:"Price,omitempty"` // 値段<br>※株式・先物・オプション銘柄の場合のみ
		Qty   float64 `json:"Qty,omitempty"`   // 数量<br>※株式・先物・オプション銘柄の場合のみ
	} `json:"Buy10,omitempty"`
	// OVER気配数量 ※株式銘柄の場合のみ
	OverSellQty float64 `json:"OverSellQty,omitempty"`
	// UNDER気配数量 ※株式銘柄の場合のみ
	UnderBuyQty float64 `json:"UnderBuyQty,omitempty"`
	// 時価総額 ※株式銘柄の場合のみ
	TotalMarketValue float64 `json:"TotalMarketValue,omitempty"`
	// 清算値 ※先物銘柄の場合のみ
	ClearingPrice float64 `json:"ClearingPrice,omitempty"`
	// インプライド・ボラティリティ ※オプション銘柄かつ日通しの場合のみ
	IV float64 `json:"IV,omitempty"`
	// ガンマ ※オプション銘柄かつ日通しの場合のみ
	Gamma float64 `json:"Gamma,omitempty"`
	// セータ ※オプション銘柄かつ日通しの場合のみ
	Theta float64 `json:"Theta,omitempty"`
	// ベガ ※オプション銘柄かつ日通しの場合のみ
	Vega float64 `json:"Vega,omitempty"`
	// デルタ ※オプション銘柄かつ日通しの場合のみ
	Delta float64 `json:"Delta,omitempty"`
	// 銘柄種別 定義値 説明 0 指数 1 現物 101 日経225先物 103 日経225OP 107 TOPIX先物 121 JPX400先物 144 NYダウ 145 日経平均VI 154 グロース250先物 155 TOPIX_REIT 171 TOPIX CORE30 901 日経平均225ミニ先物 907 TOPIXミニ先物
	SecurityType int `json:"SecurityType,omitempty"`
}

// GetInfoBoardSymbol は **GET /board/{symbol}** を呼び出して時価情報・板情報を取得します。
//
// 機能:
//   - 指定した銘柄の最新気配や約定価格、板情報を取得する。
//
// 引数:
//   - req (ReqGetInfoBoardSymbol): 銘柄コードと市場コードを含むパスパラメータ。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoBoardSymbol): 取得した時価・板情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoBoardSymbol(req ReqGetInfoBoardSymbol) (code int, res ResGetInfoBoardSymbol, err error) {
	p := "/board/{symbol}"
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

// ReqGetInfoSymbolSymbol は **GET /symbol/{symbol}** のリクエスト。
//
// 銘柄情報
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolSymbol struct {
	// 銘柄コード ※次の形式で入力してください。 [銘柄コード]@[市場コード] ※市場コードは下記の定義値から選択してください。 ※SOR市場は取扱っておりませんのでご注意ください。市場コード 定義値 説明 1 東証 3 名証 5 福証 6 札証 2 日通し 23 日中 24 夜間
	Symbol string `json:"-" path:"symbol"`
	// 追加情報出力フラグ（未指定時：true） ※追加情報は、「時価総額」、「発行済み株式数」、「決算期日」、「清算値」を意味します。 定義値 説明 true 追加情報を出力する false 追加情報を出力しない
	Addinfo string `json:"-" query:"addinfo"`
}

// ResGetInfoSymbolSymbol は **GET /symbol/{symbol}** のレスポンス。
//
// 銘柄情報
type ResGetInfoSymbolSymbol struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名
	SymbolName string `json:"SymbolName,omitempty"`
	// 銘柄略称 ※株式・先物・オプション銘柄の場合のみ
	DisplayName string `json:"DisplayName,omitempty"`
	// 市場コード ※株式・先物・オプション銘柄の場合のみ 定義値 説明 1 東証 3 名証 5 福証 6 札証 2 日通し 23 日中 24 夜間
	Exchange int `json:"Exchange,omitempty"`
	// 市場名称 ※株式・先物・オプション銘柄の場合のみ
	ExchangeName string `json:"ExchangeName,omitempty"`
	// 業種コード名 ※株式銘柄の場合のみ 定義値 説明 0050 水産・農林業 1050 鉱業 2050 建設業 3050 食料品 3100 繊維製品 3150 パルプ・紙 3200 化学 3250 医薬品 3300 石油・石炭製品 3350 ゴム製品 3400 ガラス・土石製品 3450 鉄鋼 3500 非鉄金属 3550 金属製品 3600 機械 3650 電気機器 3700 輸送用機器 3750 精密機器 3800 その他製品 4050 電気・ガス業 5050 陸運業 5100 海運業 5150 空運業 5200 倉庫・運輸関連業 5250 情報・通信業 6050 卸売業 6100 小売業 7050 銀行業 7100 証券、商品先物取引業 7150 保険業 7200 その他金融業 8050 不動産業 9050 サービス業 9999 その他
	BisCategory string `json:"BisCategory,omitempty"`
	// 時価総額 ※株式銘柄の場合のみ 追加情報出力フラグ：falseの場合、null
	TotalMarketValue float64 `json:"TotalMarketValue,omitempty"`
	// 発行済み株式数（千株） ※株式銘柄の場合のみ 追加情報出力フラグ：falseの場合、null
	TotalStocks float64 `json:"TotalStocks,omitempty"`
	// 売買単位 ※株式・先物・オプション銘柄の場合のみ
	TradingUnit float64 `json:"TradingUnit,omitempty"`
	// 決算期日 ※株式銘柄の場合のみ 追加情報出力フラグ：falseの場合、null
	FiscalYearEndBasic int `json:"FiscalYearEndBasic,omitempty"`
	// 呼値グループ ※株式・先物・オプション銘柄の場合のみ ※各呼値コードが対応する商品は以下となります。 株式の呼値の単位の詳細は [JPXページ](https://www.jpx.co.jp/equities/trading/domestic/07.html) をご覧ください。 10000：株式(通常の呼値単位の銘柄) 10003：株式(TOPIX500構成銘柄※売買単位が10口以上のETF等含む) 10004：株式(売買単位が1口のETF等) 10118 : 日経平均先物 10119 : 日経225mini 10318 : 日経平均オプション 10706 : ﾐﾆTOPIX先物 10718 : TOPIX先物 12122 : JPX日経400指数先物 14473 : NYダウ先物 14515 : 日経平均VI先物 15411 : グロース250先物 15569 : 東証REIT指数先物 17163 : TOPIXCore30指数先物 呼値コード 値段の水準 呼値単位 10000 3000円以下 1 10000 5000円以下 5 10000 30000円以下 10 10000 50000円以下 50 10000 300000円以下 100 10000 500000円以下 500 10000 3000000円以下 1000 10000 5000000円以下 5000 10000 30000000円以下 10000 10000 50000000円以下 50000 10000 50000000円超 100000 10003 1000円以下 0.1 10003 3000円以下 0.5 10003 10000円以下 1 10003 30000円以下 5 10003 100000円以下 10 10003 300000円以下 50 10003 1000000円以下 100 10003 3000000円以下 500 10003 10000000円以下 1000 10003 30000000円以下 5000 10003 30000000円超 10000 10004 10000円以下 1 10004 30000円以下 5 10004 100000円以下 10 10004 300000円以下 50 10004 1000000円以下 100 10004 3000000円以下 500 10004 10000000円以下 1000 10004 30000000円以下 5000 10004 30000000円超 10000 10118 - 10 10119 - 5 10318 100円以下 1 10318 1000円以下 5 10318 1000円超 10 10706 - 0.25 10718 - 0.5 12122 - 5 14473 - 1 14515 - 0.05 15411 - 1 15569 - 0.5 17163 - 0.5
	PriceRangeGroup string `json:"PriceRangeGroup,omitempty"`
	// 一般信用買建フラグ ※trueのとき、一般信用(長期)または一般信用(デイトレ)が買建可能 ※株式銘柄の場合のみ
	KCMarginBuy bool `json:"KCMarginBuy,omitempty"`
	// 一般信用売建フラグ ※trueのとき、一般信用(長期)または一般信用(デイトレ)が売建可能 ※株式銘柄の場合のみ
	KCMarginSell bool `json:"KCMarginSell,omitempty"`
	// 制度信用買建フラグ ※trueのとき制度信用買建可能 ※株式銘柄の場合のみ
	MarginBuy bool `json:"MarginBuy,omitempty"`
	// 制度信用売建フラグ ※trueのとき制度信用売建可能 ※株式銘柄の場合のみ
	MarginSell bool `json:"MarginSell,omitempty"`
	// 値幅上限 ※株式・先物・オプション銘柄の場合のみ
	UpperLimit float64 `json:"UpperLimit,omitempty"`
	// 値幅下限 ※株式・先物・オプション銘柄の場合のみ
	LowerLimit float64 `json:"LowerLimit,omitempty"`
	// 原資産コード ※先物・オプション銘柄の場合のみ 定義値 説明 NK225 日経225 NK300 日経300 GROWTH グロース250先物 JPX400 JPX日経400 TOPIX TOPIX NKVI 日経平均VI DJIA NYダウ TSEREITINDEX 東証REIT指数 TOPIXCORE30 TOPIX Core30
	Underlyer string `json:"Underlyer,omitempty"`
	// 限月-年月 ※「限月-年月」は「年(yyyy)/月(MM)」で表示します。 ※先物・オプション銘柄の場合のみ
	DerivMonth string `json:"DerivMonth,omitempty"`
	// 取引開始日 ※先物・オプション銘柄の場合のみ
	TradeStart int `json:"TradeStart,omitempty"`
	// 取引終了日 ※先物・オプション銘柄の場合のみ
	TradeEnd int `json:"TradeEnd,omitempty"`
	// 権利行使価格 ※オプション銘柄の場合のみ
	StrikePrice float64 `json:"StrikePrice,omitempty"`
	// プット/コール区分 ※オプション銘柄の場合のみ 定義値 説明 1 プット 2 コール
	PutOrCall int `json:"PutOrCall,omitempty"`
	// 清算値 ※先物銘柄の場合のみ 追加情報出力フラグ：falseの場合、null
	ClearingPrice float64 `json:"ClearingPrice,omitempty"`
}

// GetInfoSymbolSymbol は **GET /symbol/{symbol}** を呼び出して銘柄情報を取得します。
//
// 機能:
//   - 指定した銘柄の基本属性（銘柄種別、権利情報、価格情報など）を取得する。
//
// 引数:
//   - req (ReqGetInfoSymbolSymbol): 銘柄コードと市場コードを含むパスパラメータ。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoSymbolSymbol): 取得した銘柄情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoSymbolSymbol(req ReqGetInfoSymbolSymbol) (code int, res ResGetInfoSymbolSymbol, err error) {
	p := "/symbol/{symbol}"
	// パスパラメータの埋め込み
	p = strings.NewReplacer(
		"{symbol}", url.PathEscape(fmt.Sprint(req.Symbol)),
	).Replace(p)
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
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

// ReqGetInfoSymbolnameOption は **GET /symbolname/option** のリクエスト。
//
// オプション銘柄コード取得
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolnameOption struct {
	// オプションコード ※指定なしの場合は、日経225オプションを対象とする。 定義値 説明 NK225op 日経225オプション NK225miniop 日経225ミニオプション
	OptionCode string `json:"-" query:"OptionCode"`
	// 限月 ※限月はyyyyMM形式で指定します。0を指定した場合、直近限月となります。 ※取引最終日に「0」（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	DerivMonth string `json:"-" query:"DerivMonth" required:"true"`
	// コール or プット ※大文字小文字は区別しません。 定義値 説明 P PUT C CALL
	PutOrCall string `json:"-" query:"PutOrCall" required:"true"`
	// 権利行使価格 ※0を指定した場合、APIを実行した時点でのATMとなります。
	StrikePrice int `json:"-" query:"StrikePrice" required:"true"`
}

// ResGetInfoSymbolnameOption は **GET /symbolname/option** のレスポンス。
//
// オプション銘柄コード取得
type ResGetInfoSymbolnameOption struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名称
	SymbolName string `json:"SymbolName,omitempty"`
}

// GetInfoSymbolnameOption は **GET /symbolname/option** を呼び出してオプション銘柄コードを取得します。
//
// 機能:
//   - オプション銘柄コード区分や限月、権利行使価格などの条件に基づき、取引可能なオプション銘柄コード一覧を取得する。
//
// 引数:
//   - req (ReqGetInfoSymbolnameOption): オプションコード区分や限月指定などの取得条件。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoSymbolnameOption): 取得したオプション銘柄コード情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoSymbolnameOption(req ReqGetInfoSymbolnameOption) (code int, res ResGetInfoSymbolnameOption, err error) {
	p := "/symbolname/option"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.OptionCode != "" {
		v.Set("OptionCode", fmt.Sprint(req.OptionCode))
	}
	v.Set("DerivMonth", fmt.Sprint(req.DerivMonth))
	if req.PutOrCall != "" {
		v.Set("PutOrCall", fmt.Sprint(req.PutOrCall))
	}
	if req.StrikePrice != 0 {
		v.Set("StrikePrice", fmt.Sprint(req.StrikePrice))
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

// ReqGetInfoSymbolnameMinioptionweekly は **GET /symbolname/minioptionweekly** のリクエスト。
//
// ミニオプション（限週）銘柄コード取得
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoSymbolnameMinioptionweekly struct {
	// 限月 ※限月はyyyyMM形式で指定します。0を指定した場合、直近限月となります。 ※取引最終日に「0」（直近限月）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限月の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	DerivMonth string `json:"-" query:"DerivMonth" required:"true"`
	// 限週 ※限週は0,1,3,4,5のいずれかを指定します。0を指定した場合、指定した限月の直近限週となります。 ※取引最終日に「0」（直近限週）を指定した場合、日中・夜間の時間帯に関わらず、取引最終日を迎える限週の銘柄コードを返します。取引最終日を迎える銘柄の取引は日中取引をもって終了となりますので、ご注意ください。
	DerivWeekly int `json:"-" query:"DerivWeekly" required:"true"`
	// コール or プット ※大文字小文字は区別しません。 定義値 説明 P PUT C CALL
	PutOrCall string `json:"-" query:"PutOrCall" required:"true"`
	// 権利行使価格 ※0を指定した場合、APIを実行した時点でのATMとなります。
	StrikePrice int `json:"-" query:"StrikePrice" required:"true"`
}

// ResGetInfoSymbolnameMinioptionweekly は **GET /symbolname/minioptionweekly** のレスポンス。
//
// ミニオプション（限週）銘柄コード取得
type ResGetInfoSymbolnameMinioptionweekly struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 銘柄名称
	SymbolName string `json:"SymbolName,omitempty"`
}

// GetInfoSymbolnameMinioptionweekly は **GET /symbolname/minioptionweekly** を呼び出してミニオプション（限週）の銘柄コードを取得します。
//
// 機能:
//   - ミニオプションの限月・限週条件を指定し、取引可能な銘柄コード一覧を取得する。
//
// 引数:
//   - req (ReqGetInfoSymbolnameMinioptionweekly): 限月や限週を指定するリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoSymbolnameMinioptionweekly): 取得したミニオプション銘柄コード情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoSymbolnameMinioptionweekly(req ReqGetInfoSymbolnameMinioptionweekly) (code int, res ResGetInfoSymbolnameMinioptionweekly, err error) {
	p := "/symbolname/minioptionweekly"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	v.Set("DerivMonth", fmt.Sprint(req.DerivMonth))
	if req.DerivWeekly != 0 {
		v.Set("DerivWeekly", fmt.Sprint(req.DerivWeekly))
	}
	if req.PutOrCall != "" {
		v.Set("PutOrCall", fmt.Sprint(req.PutOrCall))
	}
	if req.StrikePrice != 0 {
		v.Set("StrikePrice", fmt.Sprint(req.StrikePrice))
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

// ReqGetInfoRanking は **GET /ranking** のリクエスト。
//
// 詳細ランキング
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoRanking struct {
	// 種別 ※信用情報ランキングに「福証」「札証」を指定した場合は、空レスポンスになります 定義値 説明 1 値上がり率（デフォルト） 2 値下がり率 3 売買高上位 4 売買代金 5 TICK回数 6 売買高急増 7 売買代金急増 8 信用売残増 9 信用売残減 10 信用買残増 11 信用買残減 12 信用高倍率 13 信用低倍率 14 業種別値上がり率 15 業種別値下がり率
	Type string `json:"-" query:"Type" required:"true"`
	// 市場 ※業種別値上がり率・値下がり率に市場を指定しても無視されます 定義値 説明 ALL 全市場（デフォルト） T 東証全体 TP 東証プライム TS 東証スタンダード TG グロース250 M 名証 FK 福証 S 札証
	ExchangeDivision string `json:"-" query:"ExchangeDivision" required:"true"`
}

// ResGetInfoRanking は **GET /ranking** のレスポンス。
//
// 詳細ランキング
type ResGetInfoRanking struct {
}

// GetInfoRanking は **GET /ranking** を呼び出して各種ランキング情報を取得します。
//
// 機能:
//   - 詳細ランキング画面と同様に、出来高や値上がり率などのランキングを取得する。
//
// 引数:
//   - req (ReqGetInfoRanking): 取得するランキングの種別などを指定するリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoRanking): 取得したランキング情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoRanking(req ReqGetInfoRanking) (code int, res ResGetInfoRanking, err error) {
	p := "/ranking"
	// クエリパラメータの構築（zero値は送信しません）
	v := url.Values{}
	if req.Type != "" {
		v.Set("Type", fmt.Sprint(req.Type))
	}
	if req.ExchangeDivision != "" {
		v.Set("ExchangeDivision", fmt.Sprint(req.ExchangeDivision))
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

// ReqGetInfoExchangeSymbol は **GET /exchange/{symbol}** のリクエスト。
//
// 為替情報
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoExchangeSymbol struct {
	// 通貨 定義値 内容 usdjpy USD/JPY eurjpy EUR/JPY gbpjpy GBP/JPY audjpy AUD/JPY chfjpy CHF/JPY cadjpy CAD/JPY nzdjpy NZD/JPY zarjpy ZAR/JPY eurusd EUR/USD gbpusd GBP/USD audusd AUD/USD
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoExchangeSymbol は **GET /exchange/{symbol}** のレスポンス。
//
// 為替情報
type ResGetInfoExchangeSymbol struct {
	// 通貨
	Symbol string `json:"Symbol,omitempty"`
	// BID
	BidPrice float64 `json:"BidPrice,omitempty"`
	// SP
	Spread float64 `json:"Spread,omitempty"`
	// ASK
	AskPrice float64 `json:"AskPrice,omitempty"`
	// 前日比
	Change float64 `json:"Change,omitempty"`
	// 時刻 ※HH:mm:ss形式
	Time string `json:"Time,omitempty"`
}

// GetInfoExchangeSymbol は **GET /exchange/{symbol}** を呼び出して為替情報を取得します。
//
// 機能:
//   - 指定した通貨ペアの気配値や現在値など、マネービューに相当する為替情報を取得する。
//
// 引数:
//   - req (ReqGetInfoExchangeSymbol): 通貨ペアと市場コードをパスパラメータで指定するリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoExchangeSymbol): 取得した為替情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoExchangeSymbol(req ReqGetInfoExchangeSymbol) (code int, res ResGetInfoExchangeSymbol, err error) {
	p := "/exchange/{symbol}"
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

// ReqGetInfoRegulationsSymbol は **GET /regulations/{symbol}** のリクエスト。
//
// 規制情報
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoRegulationsSymbol struct {
	// 銘柄コード ※次の形式で入力してください。 [銘柄コード]@[市場コード] ※市場コードは下記の定義値から選択してください。 市場コード 定義値 説明 1 東証 3 名証 5 福証 6 札証
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoRegulationsSymbol は **GET /regulations/{symbol}** のレスポンス。
//
// 規制情報
type ResGetInfoRegulationsSymbol struct {
	// 銘柄コード ※対象商品は、株式のみ
	Symbol string `json:"Symbol,omitempty"`
	// 規制情報
	RegulationsInfo []struct {
		Exchange      int    `json:"Exchange,omitempty"`      // 規制市場 <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>全対象</td> </tr> <tr> <td>1</td> <td>東証</td> </tr> <tr> <td>3</td> <td>名証</td> </tr> <tr> <td>5</td> <td>福証</td> </tr> <tr> <td>6</td> <td>札証</td> </tr> <tr> <td>9</td> <td>SOR</td> </tr> <tr> <td>10</td> <td>CXJ</td> </tr> <tr> <td>21</td> <td>JNX</td> </tr> </tbody> </table>
		Product       int    `json:"Product,omitempty"`       // 規制取引区分<br> ※空売り規制の場合、「4：新規」 <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>全対象</td> </tr> <tr> <td>1</td> <td>現物</td> </tr> <tr> <td>2</td> <td>信用新規（制度）</td> </tr> <tr> <td>3</td> <td>信用新規（一般）</td> </tr> <tr> <td>4</td> <td>新規</td> </tr> <tr> <td>5</td> <td>信用返済（制度）</td> </tr> <tr> <td>6</td> <td>信用返済（一般）</td> </tr> <tr> <td>7</td> <td>返済</td> </tr> <tr> <td>8</td> <td>品受</td> </tr> <tr> <td>9</td> <td>品渡</td> </tr> </tbody> </table>
		Side          string `json:"Side,omitempty"`          // 規制売買<br> ※空売り規制の場合、「1：売」 <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>0</td> <td>全対象</td> </tr> <tr> <td>1</td> <td>売</td> </tr> <tr> <td>2</td> <td>買</td> </tr> </tbody> </table>
		Reason        string `json:"Reason,omitempty"`        // 理由<br>※空売り規制の場合、「空売り規制」
		LimitStartDay string `json:"LimitStartDay,omitempty"` // 制限開始日<br>yyyy/MM/dd HH:mm形式 <br>※空売り規制の場合、null
		LimitEndDay   string `json:"LimitEndDay,omitempty"`   // 制限終了日<br>yyyy/MM/dd HH:mm形式 <br>※空売り規制の場合、null
		Level         int    `json:"Level,omitempty"`         // コンプライアンスレベル<br> ※空売り規制の場合、null <table> <thead> <tr> <th>定義値</th> <th>内容</th> </tr> </thead> <tbody> <tr> <td>０</td> <td>規制無し</td> </tr> <tr> <td>１</td> <td>ワーニング</td> </tr> <tr> <td>２</td> <td>エラー</td> </tr> </tbody> </table>
	} `json:"RegulationsInfo,omitempty"`
}

// GetInfoRegulationsSymbol は **GET /regulations/{symbol}** を呼び出して規制情報を取得します。
//
// 機能:
//   - 指定した銘柄に適用される各種売買規制や空売り規制の情報を取得する。
//
// 引数:
//   - req (ReqGetInfoRegulationsSymbol): 銘柄コードと市場コードを含むパスパラメータ。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoRegulationsSymbol): 取得した規制情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoRegulationsSymbol(req ReqGetInfoRegulationsSymbol) (code int, res ResGetInfoRegulationsSymbol, err error) {
	p := "/regulations/{symbol}"
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

// ReqGetInfoPrimaryexchangeSymbol は **GET /primaryexchange/{symbol}** のリクエスト。
//
// 優先市場
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoPrimaryexchangeSymbol struct {
	// 銘柄コード
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoPrimaryexchangeSymbol は **GET /primaryexchange/{symbol}** のレスポンス。
//
// 優先市場
type ResGetInfoPrimaryexchangeSymbol struct {
	// 銘柄コード ※対象商品は、株式のみ
	Symbol string `json:"Symbol,omitempty"`
	// 優先市場 定義値 説明 1 東証 3 名証 5 福証 6 札証
	PrimaryExchange int `json:"PrimaryExchange,omitempty"`
}

// GetInfoPrimaryexchangeSymbol は **GET /primaryexchange/{symbol}** を呼び出して株式の優先市場を取得します。
//
// 機能:
//   - 指定した株式銘柄について、優先する市場コードを取得する。
//
// 引数:
//   - req (ReqGetInfoPrimaryexchangeSymbol): 銘柄コードを含むパスパラメータ。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoPrimaryexchangeSymbol): 優先市場情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoPrimaryexchangeSymbol(req ReqGetInfoPrimaryexchangeSymbol) (code int, res ResGetInfoPrimaryexchangeSymbol, err error) {
	p := "/primaryexchange/{symbol}"
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

// ReqGetInfoApisoftlimit は **GET /apisoftlimit** のリクエスト。
//
// ソフトリミット
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoApisoftlimit struct {
}

// ResGetInfoApisoftlimit は **GET /apisoftlimit** のレスポンス。
//
// ソフトリミット
type ResGetInfoApisoftlimit struct {
	// 現物のワンショット上限 ※単位は万円
	Stock float64 `json:"Stock,omitempty"`
	// 信用のワンショット上限 ※単位は万円
	Margin float64 `json:"Margin,omitempty"`
	// 先物のワンショット上限 ※単位は枚
	Future float64 `json:"Future,omitempty"`
	// ミニ先物のワンショット上限 ※単位は枚
	FutureMini float64 `json:"FutureMini,omitempty"`
	// マイクロ先物のワンショット上限 ※単位は枚
	FutureMicro float64 `json:"FutureMicro,omitempty"`
	// オプションのワンショット上限 ※単位は枚
	Option float64 `json:"Option,omitempty"`
	// ミニオプションのワンショット上限 ※単位は枚
	MiniOption float64 `json:"MiniOption,omitempty"`
	// kabuステーションのバージョン
	KabuSVersion string `json:"KabuSVersion,omitempty"`
}

// GetInfoApisoftlimit は **GET /apisoftlimit** を呼び出してソフトリミット値を取得します。
//
// 機能:
//   - kabuステーション API に設定された各種ソフトリミット値を取得する。
//
// 引数:
//   - req (ReqGetInfoApisoftlimit): 空のリクエスト。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoApisoftlimit): ソフトリミットの値。
//   - err (error): 通信エラーまたは APIError。
func GetInfoApisoftlimit(req ReqGetInfoApisoftlimit) (code int, res ResGetInfoApisoftlimit, err error) {
	p := "/apisoftlimit"
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

// ReqGetInfoMarginMarginpremiumSymbol は **GET /margin/marginpremium/{symbol}** のリクエスト。
//
// プレミアム料取得
//
//   - 認証: 必須（X-API-KEY ヘッダー）
type ReqGetInfoMarginMarginpremiumSymbol struct {
	// 銘柄コード
	Symbol string `json:"-" path:"symbol"`
}

// ResGetInfoMarginMarginpremiumSymbol は **GET /margin/marginpremium/{symbol}** のレスポンス。
//
// プレミアム料取得
type ResGetInfoMarginMarginpremiumSymbol struct {
	// 銘柄コード
	Symbol string `json:"Symbol,omitempty"`
	// 一般信用（長期）
	GeneralMargin struct {
		MarginPremiumType  int     `json:"MarginPremiumType,omitempty"`  // プレミアム料入力区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>null</td> <td>一般信用（長期）非対応銘柄</td> </tr> <tr> <td>0</td> <td>プレミアム料がない銘柄</td> </tr> <tr> <td>1</td> <td>プレミアム料が固定の銘柄</td> </tr> <tr> <td>2</td> <td>プレミアム料が入札で決定する銘柄</td> </tr> </tbody> </table>
		MarginPremium      float64 `json:"MarginPremium,omitempty"`      // 確定プレミアム料<br> ※入札銘柄の場合、入札受付中は随時更新します。受付時間外は、確定したプレミアム料を返します。<br> ※非入札銘柄の場合、常に固定値を返します。<br> ※信用取引不可の場合、nullを返します。<br> ※19:30~翌営業日のプレミアム料になります。
		UpperMarginPremium float64 `json:"UpperMarginPremium,omitempty"` // 上限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		LowerMarginPremium float64 `json:"LowerMarginPremium,omitempty"` // 下限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		TickMarginPremium  float64 `json:"TickMarginPremium,omitempty"`  // プレミアム料刻値<br> ※入札可能銘柄以外は、nullを返します。
	} `json:"GeneralMargin,omitempty"`
	// 一般信用（デイトレ）
	DayTrade struct {
		MarginPremiumType  int     `json:"MarginPremiumType,omitempty"`  // プレミアム料入力区分 <table> <thead> <tr> <th>定義値</th> <th>説明</th> </tr> </thead> <tbody> <tr> <td>null</td> <td>一般信用（デイトレ）非対応銘柄</td> </tr> <tr> <td>0</td> <td>プレミアム料がない銘柄</td> </tr> <tr> <td>1</td> <td>プレミアム料が固定の銘柄</td> </tr> <tr> <td>2</td> <td>プレミアム料が入札で決定する銘柄</td> </tr> </tbody> </table>
		MarginPremium      float64 `json:"MarginPremium,omitempty"`      // 確定プレミアム料<br> ※入札銘柄の場合、入札受付中は随時更新します。受付時間外は、確定したプレミアム料を返します。<br> ※非入札銘柄の場合、常に固定値を返します。<br> ※信用取引不可の場合、nullを返します。<br> ※19:30~翌営業日のプレミアム料になります。
		UpperMarginPremium float64 `json:"UpperMarginPremium,omitempty"` // 上限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		LowerMarginPremium float64 `json:"LowerMarginPremium,omitempty"` // 下限プレミアム料<br> ※プレミアム料がない場合は、nullを返します。
		TickMarginPremium  float64 `json:"TickMarginPremium,omitempty"`  // プレミアム料刻値<br> ※入札可能銘柄以外は、nullを返します。
	} `json:"DayTrade,omitempty"`
}

// GetInfoMarginMarginpremiumSymbol は **GET /margin/marginpremium/{symbol}** を呼び出してプレミアム料を取得します。
//
// 機能:
//   - 指定した銘柄のプレミアム料や日計りに関する情報を取得する。
//
// 引数:
//   - req (ReqGetInfoMarginMarginpremiumSymbol): 銘柄コードと市場コードを含むパスパラメータ。
//
// 返り値:
//   - code (int): HTTP ステータスコード。
//   - res (ResGetInfoMarginMarginpremiumSymbol): プレミアム料情報。
//   - err (error): 通信エラーまたは APIError。
func GetInfoMarginMarginpremiumSymbol(req ReqGetInfoMarginMarginpremiumSymbol) (code int, res ResGetInfoMarginMarginpremiumSymbol, err error) {
	p := "/margin/marginpremium/{symbol}"
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
