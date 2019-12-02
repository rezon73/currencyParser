package exchange

import (
	"currencyParser/cache"
	"currencyParser/entity"
	"currencyParser/service/config"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type BinanceExchange struct {
	Config       *config.Config
	MainDatabase *gorm.DB
}

type getBinanceQuoteResponse struct {
	Code    int     `json:"code,omitempty"`
	Message string  `json:"msg,omitempty"`
	Price   string  `json:"price,omitempty"`
	Symbol   string  `json:"symbol,omitempty"`
}

func (exchange BinanceExchange) GetExchangeId() int {
	return entity.EXCHANGE_ID_BINANCE
}

func (exchange BinanceExchange) GetExchangePrice(symbol string) (float64, error) {
	url := exchange.Config.Binance.GetQuoteUrl + "?symbol=" + symbol
	cacheKey := "getQuote" + url
	body, cacheErr := cache.Get(cacheKey)

	if cacheErr != nil || len(body) == 0 {
		request, _ := http.NewRequest("GET", url, nil)
		request.Header.Add("cache-control", "no-cache")
		request.Header.Add("X-MBX-APIKEY", exchange.Config.Binance.ApiKey)

		responseRaw, _ := http.DefaultClient.Do(request)

		defer responseRaw.Body.Close()
		body, _ = ioutil.ReadAll(responseRaw.Body)

		cacheErr = cache.Set(cacheKey, body, 5)
		if cacheErr != nil {
			log.Println(cacheErr)
		}
	}

	var response getBinanceQuoteResponse

	err := json.Unmarshal(body, &response)

	if err == nil {
		if response.Code < 0 {
			err = errors.New(response.Message)
		}
	}

	if err != nil {
		return 0.0, err
	}

	price, err := strconv.ParseFloat(response.Price, 64)

	return price, err
}