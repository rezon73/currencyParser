package controller

import (
	"currencyParser/entity"
	"currencyParser/service/mainDatabase"
	"encoding/json"
	"net/http"
)

type IndexController struct {}

func (controller IndexController) GetQuoteHandler(writer http.ResponseWriter, request *http.Request) {
	symbolNames := request.FormValue("symbols")

	var symbols []entity.Symbol
	var actualQuote entity.ActualQuote
	prices := make(map[string]float64)

	mainDatabase.GetInstance(0).Find(&symbols, "name in (?) and is_deleted = false and parent_id = 0", symbolNames)
	for _, symbol := range symbols {
		mainDatabase.GetInstance(0).Model(&symbol).Related(&actualQuote, "ActualQuote")
		if actualQuote.Id == 0 {
			continue
		}

		prices[symbol.Name] = actualQuote.Price
	}

	bytes, _ := json.Marshal(prices)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}

func (controller IndexController) GetSymbolsHandler(writer http.ResponseWriter, request *http.Request) {
	var symbols []entity.Symbol
	var symbolNames []string

	mainDatabase.GetInstance(0).Find(&symbols, "is_deleted = false and parent_id = 0")
	for _, symbol := range symbols {
		symbolNames = append(symbolNames, symbol.Name)
	}

	bytes, _ := json.Marshal(symbolNames)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}