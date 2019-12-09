package config

type Log struct {
	PushUrl string `envconfig:"LOG_PUSH_URL"`
	Source  string `envconfig:"LOG_SOURCE"`
}
