package config

type Binance struct {
	ApiKey      string `envconfig:"BINANCE_KEY"`
	GetQuoteUrl string `envconfig:"BINANCE_GET_QUOTE_URL"`
}
