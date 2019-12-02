package quoteUpdater

import (
	"currencyParser/entity"
	"currencyParser/service/config"
	"currencyParser/service/mainDatabase"
	"currencyParser/service/quoteUpdater/exchange"
	"github.com/jinzhu/gorm"
)

type QuoteUpdater struct {
	Config       *config.Config
	MainDatabase *gorm.DB
	Exchange     exchange.Exchange
}

func (quoteUpdater *QuoteUpdater) GetActualQuote(symbol string) (float64, error) {
	return 0.0, nil
}

func (quoteUpdater *QuoteUpdater) UpdateActualQuote(symbol entity.Symbol) error {
	var actualQuote entity.ActualQuote
	exchangePrice, err := quoteUpdater.Exchange.GetExchangePrice(symbol.GetSymbolNameByExchange(quoteUpdater.Exchange.GetExchangeId()))

	if err != nil {
		return err
	}

	mainDatabase.GetInstance(0).First(
		&actualQuote,
		"symbol_id = ? and exchange_id = ?",
		symbol.Id,
		quoteUpdater.Exchange.GetExchangeId(),
	)

	if actualQuote.Id > 0 {
		mainDatabase.GetInstance(0).Model(&actualQuote).Update("Price", exchangePrice).Where("id = ?", actualQuote.Id)
	} else {
		mainDatabase.GetInstance(0).Create(&entity.ActualQuote{
			Price: exchangePrice,
			ExchangeId: quoteUpdater.Exchange.GetExchangeId(),
			SymbolId: symbol.Id,
		})
	}

	return err
}

func (quoteUpdater *QuoteUpdater) UpdateActualQuotesExceptSeveral(exceptSymbols []entity.Symbol) error {
	var symbols []entity.Symbol
	var err error

	var exceptSymbolIds []int
	for _, exceptSymbol := range exceptSymbols {
		exceptSymbolIds = append(exceptSymbolIds, exceptSymbol.Id)
	}
	query := quoteUpdater.MainDatabase.Where("id not in (?) and parent_id = 0 and is_deleted = false", exceptSymbolIds)
	query.Find(&symbols)

	for _, symbol := range symbols {
		err = quoteUpdater.UpdateActualQuote(symbol)
		if err != nil {
			return err
		}
	}

	return err
}

func (quoteUpdater *QuoteUpdater) UpdateActualQuotes() error {
	var symbols []entity.Symbol
	var err error

	quoteUpdater.MainDatabase.Find(&symbols, "parent_id = 0 and is_deleted = false")

	for _, symbol := range symbols {
		err = quoteUpdater.UpdateActualQuote(symbol)
		if err != nil {
			return err
		}
	}

	return err
}
