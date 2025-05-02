package model

type TokenInfo struct {
	Name      string  `json:"token_name"`
	Symbol    string  `json:"symbol"`
	PriceUSD  float64 `json:"price_usd"`
	MarketCap float64 `json:"market_cap_usd"`
}

var MockTokenDB = map[string]TokenInfo{
	"BTC": {
		Name:      "Bitcoin",
		Symbol:    "BTC",
		PriceUSD:  65000.00,
		MarketCap: 1200000000000,
	},
	"ETH": {
		Name:      "Ethereum",
		Symbol:    "ETH",
		PriceUSD:  3000.00,
		MarketCap: 450000000000,
	},
	"SOL": {
		Name:      "Solana",
		Symbol:    "SOL",
		PriceUSD:  150.00,
		MarketCap: 65000000000,
	},
}
