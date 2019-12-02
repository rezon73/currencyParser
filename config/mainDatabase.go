package config

type MainDatabase struct {
	DSN string `envconfig:"MYSQL_DSN"` // %s:%s@tcp(%s:%d)/%s?charset=utf8
}
