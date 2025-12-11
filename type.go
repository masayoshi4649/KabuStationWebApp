package main

// Config は、システム設定とトレード設定をまとめて保持するルート構造体です。
type Config struct {
	Trade struct {
		FutureCode string  `toml:"FUTURE_CODE"`
		DerivMonth string  `toml:"DERIV_MONTH"`
		OneTick    float64 `toml:"ONE_TICK"`
		Exchange   int     `toml:"EXCHANGE"`
	} `toml:"TRADE"`
}

// BookRow は、板表示に利用する 1 行分の価格と数量を表現します。
type BookRow struct {
	Price   float64 `json:"Price"`
	Current bool    `json:"Current"`
	SellQty float64 `json:"SellQty"`
	BuyQty  float64 `json:"BuyQty"`
}
