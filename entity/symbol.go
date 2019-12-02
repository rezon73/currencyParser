package entity

import (
	"currencyParser/service/mainDatabase"
)

type Symbol struct {
	Id         int      `gorm:"AUTO_INCREMENT;not null"`
	Name       string   `gorm:"not_null;unique_index:user_symbol_index"`
	ParentId   int      `gorm:"default:0"`
	ExchangeId int      `gorm:"default:0"`
	IsDeleted  bool     `gorm:"default:'false'"`
	ActualQuote ActualQuote `gorm:"foreignkey:SymbolId"`
}

func (symbol *Symbol) GetSymbolNameByExchange(exchangeId int) string {
	var aliasSymbols []Symbol
	mainDatabase.GetInstance(0).Find(&aliasSymbols, "parent_id = ? and exchange_id = ?", symbol.Id, exchangeId)
	if len(aliasSymbols) > 1 {
		panic("Too much aliases for symbol " + symbol.Name + " and exchange " + string(exchangeId))
	}

	if len(aliasSymbols) == 1 {
		return aliasSymbols[0].Name
	}

	return symbol.Name
}