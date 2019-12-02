package exchange

type Exchange interface {
	GetExchangeId()                      int
	GetExchangePrice(symbolName string)  (float64, error)
}
