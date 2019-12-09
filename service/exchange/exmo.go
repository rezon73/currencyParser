package exchange

import (
	"currencyParser/cache"
	"currencyParser/entity"
	"currencyParser/service/config"
	"currencyParser/service/logService"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ExmoExchange struct {
	Config       *config.Config
	MainDatabase *gorm.DB
}

type getExmoQuoteResponse struct {
	Price   string  `json:"sell_price,omitempty"`
}

func (exchange ExmoExchange) GetExchangeId() int {
	return entity.EXCHANGE_ID_EXMO
}

func (exchange ExmoExchange) GetExchangePrice(symbol string) (float64, error) {
	url := exchange.Config.Binance.GetQuoteUrl + "?symbol=" + symbol
	cacheKey := "getQuote" + url
	body, cacheErr := cache.Get(cacheKey)

	if cacheErr != nil || len(body) == 0 {
		request, _ := http.NewRequest("GET", exchange.Config.Exmo.GetQuoteUrl, nil)
		request.Header.Add("cache-control", "no-cache")
		responseRaw, _ := http.DefaultClient.Do(request)

		defer responseRaw.Body.Close()
		body, _ = ioutil.ReadAll(responseRaw.Body)

		cacheErr = cache.Set(cacheKey, body, 5)
		if cacheErr != nil {
			logService.Warn(cacheErr)
		}
	}

	var response map[string]getExmoQuoteResponse

	err := json.Unmarshal(body, &response)
	if err == nil {
		if _, isExist := response[symbol]; !isExist {
			err = errors.New("Invalid symbol " + symbol)
		}
	}

	if err != nil {
		return 0.0, err
	}

	price, err := strconv.ParseFloat(response[symbol].Price, 64)

	return price, err
}