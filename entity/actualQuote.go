package entity

import (
	"currencyParser/service/mainDatabase"
	"time"
)

type ActualQuote struct {
	Id          int       `gorm:"AUTO_INCREMENT;not null"`
	SymbolId    int       `gorm:"index:symbol_id_index"`
	ExchangeId  int       `gorm:"index:exchange_id_index"`
	Price       float64   `gorm:"not_null"`
	UpdatedAt   *time.Time
}

const EXCHANGE_ID_BINANCE int = 1;
const EXCHANGE_ID_EXMO    int = 2;

func (quote *ActualQuote) GetSymbol() Symbol {
	var symbol Symbol
	mainDatabase.GetInstance(0).Find(&symbol, "id = ?", quote.SymbolId)

	return symbol
}