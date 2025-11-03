package main

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

// 表示用1行
type BookRow struct {
	Price   float64 `json:"Price"`
	Current bool    `json:"Current"`
	SellQty float64 `json:"SellQty"`
	BuyQty  float64 `json:"BuyQty"`
}
