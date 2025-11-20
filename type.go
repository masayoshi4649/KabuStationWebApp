package main

// Config はシステム設定とトレード設定を保持するルート構造体。
type Config struct {
	System struct {
		EnvName string `toml:"ENV_NAME"`
	} `toml:"SYSTEM"`
	Trade struct {
		FutureCode string  `toml:"FUTURE_CODE"`
		DerivMonth string  `toml:"DERIV_MONTH"`
		OneTick    float64 `toml:"ONE_TICK"`
		Exchange   int     `toml:"EXCHANGE"`
	} `toml:"TRADE"`
}

// BookRow は描画用に整形した1行分の板情報を表す。
type BookRow struct {
	Price   float64 `json:"Price"`
	Current bool    `json:"Current"`
	SellQty float64 `json:"SellQty"`
	BuyQty  float64 `json:"BuyQty"`
}
