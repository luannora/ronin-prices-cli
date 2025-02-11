package prices

import "net/http"

type PricesInput struct {
	Source      string
	APIToken    string
	HTTPClient  *http.Client
	APIEndpoint string
}

type RoninPricesResp []CurrencyConvert

type CurrencyConvert struct {
	Convert   string  `json:"convert"`
	Quotes    []Quote `json:"quotes"`
	Source    string  `json:"source"`
	UpdatedAt int     `json:"updated_at"`
}

type Quote struct {
	Convert string `json:"convert"`
	Price   string `json:"price"`
	Symbol  string `json:"symbol"`
}

type Sources struct {
	Sources Source `json:"sources"`
}

type Source struct {
	Binance        bool `json:"binance"`
	Blue           bool `json:"blue"`
	CoinMarketCap  bool `json:"coinmarketcap"`
	Stratum  	   bool `json:"stratum"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
