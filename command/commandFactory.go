package command

import (
	"currencyParser/service/config"
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
			panic("Factory: set correct command")
	}
}
