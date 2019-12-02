package config

type Cache struct {
	Servers string `envconfig:"MEMCACHE_SERVERS"`
}
