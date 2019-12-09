package command

import (
	"currencyParser/entity"
	"currencyParser/service/config"
	"currencyParser/service/exchange"
	"currencyParser/service/logService"
	"currencyParser/service/mainDatabase"
	"currencyParser/service/quoteUpdater"
	"flag"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

type UpdateQuote struct {
	Config           *config.Config
	MainDatabase     *gorm.DB
}

var updateQuoteExchangeId int
var updateQuoteSymbols string
var updateQuoteExceptSymbolMode bool

func (command UpdateQuote) Exec() error {
	var err error

	if !flag.Parsed() {
		exchangeId := flag.Int("exchange", 0, "exchange id (required)")
		symbolFlag := flag.String("symbols", "", "symbol name (optional)")
		flag.Parse()

		updateQuoteExchangeId = *exchangeId
		updateQuoteSymbols = strings.ReplaceAll(*symbolFlag, "^", "")
		updateQuoteExceptSymbolMode = strings.Index(*symbolFlag, "^") != -1
	}

	if updateQuoteExchangeId == 0 {
		logService.Fatal("UpdateQuote: set exchange [-exchange=...]")
	}

	var symbols []entity.Symbol
	symbolNames := strings.Split(updateQuoteSymbols, ",")
	if len(symbolNames) > 0 {
		command.MainDatabase.Where("name in (?) and is_deleted = false and parent_id = 0", symbolNames).Find(&symbols)
	}

	updater := quoteUpdater.QuoteUpdater{
		Config:       config.GetConfig(),
		MainDatabase: mainDatabase.GetInstance(0),
		Exchange:     command.spawnExchange(updateQuoteExchangeId),
	}

	if len(symbols) > 0 {
		if updateQuoteExceptSymbolMode {
			err = updater.UpdateActualQuotesExceptSeveral(symbols)
		} else {
			for _, symbol := range symbols {

				err = updater.UpdateActualQuote(symbol)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	} else {
		err = updater.UpdateActualQuotes()
	}

	return err
}

func (command UpdateQuote) spawnExchange(exchangeId int) exchange.Exchange {
	switch exchangeId {
		case entity.EXCHANGE_ID_BINANCE:
			return exchange.BinanceExchange{
				Config:       command.Config,
				MainDatabase: command.MainDatabase,
			}
		case entity.EXCHANGE_ID_EXMO:
			return exchange.ExmoExchange{
				Config:       command.Config,
				MainDatabase: command.MainDatabase,
			}
		default:
			logService.Fatal("Set correct exchange ID")
			return nil
	}
}