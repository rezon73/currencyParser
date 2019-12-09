package command

import (
	"currencyParser/service/config"
	"currencyParser/service/logService"
	"github.com/jinzhu/gorm"
)

type Factory struct {
	CommandName  string
	Config       *config.Config
	MainDatabase *gorm.DB
}

func (factory Factory) CreateCommand() Command {
	switch factory.CommandName {
		case "updateQuote":
			return UpdateQuote{
				Config: factory.Config,
				MainDatabase: factory.MainDatabase,
			}
		default:
			logService.Fatal("Factory: set correct command")
			return nil
	}
}
