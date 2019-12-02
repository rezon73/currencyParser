package cache

import (
	"currencyParser/service/config"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
)

var cacheInstance *memcache.Client

func init() {
	InitInstance(config.GetConfig().Cache.Servers)
}

func InitInstance(instances string) {
	hosts := strings.Split(instances, ",")

	cacheInstance = memcache.New(hosts...)
}

func Get(key string) ([]byte, error) {
	item, err := cacheInstance.Get(key)
	if err == nil {
		return item.Value, nil
	}

	return nil, err
}

func Set(key string, value []byte, expiration int32) error {
	return cacheInstance.Set(&memcache.Item{Key: key, Value: value, Expiration: expiration})
}