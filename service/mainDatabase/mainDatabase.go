package mainDatabase

import (
	"currencyParser/service/config"
	"currencyParser/service/logService"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instances map[int]*gorm.DB = make(map[int]*gorm.DB, 1)

func init() {
	InitInstance(0, config.GetConfig().MainDatabase.DSN)
}

func InitInstance(instanceId int, dsn string) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logService.Error("failed to connect database")
	}
	db.SingularTable(true)

	instances[instanceId] = db
}

func GetInstance(instanceId int) *gorm.DB {
	return instances[instanceId]
}

func Close() {
	for _, instance := range instances {
		instance.Close()
	}
}