package config

type Telegram struct {
	Enabled bool   `envconfig:"TELEGRAM_ENABLED"`
	Token   string `envconfig:"TELEGRAM_TOKEN"`
	ChatId  int64  `envconfig:"TELEGRAM_CHAT_ID"`
}
