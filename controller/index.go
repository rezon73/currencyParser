package controller

import (
	"currencyParser/entity"
	"currencyParser/service/logService"
	"currencyParser/service/mainDatabase"
	"encoding/json"
	"net/http"
	"strings"
)

type IndexController struct {}

func (controller IndexController) GetQuoteHandler(writer http.ResponseWriter, request *http.Request) {
	symbolNames := strings.Split(request.FormValue("symbols"), ",")

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

	bytes, err := json.Marshal(prices)
	if err != nil {
		logService.Error(err)
	}

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

	bytes, err := json.Marshal(symbolNames)
	if err != nil {
		logService.Error(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}