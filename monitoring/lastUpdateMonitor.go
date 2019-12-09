package monitoring

import (
	"currencyParser/entity"
	"currencyParser/service/config"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type LastUpdateMonitor struct {
	Config           *config.Config
	MainDatabase     *gorm.DB
	DelayErrorLevel  int
}

func (monitor LastUpdateMonitor) Check() []error {
	var errs []error

	var quotes []entity.ActualQuote
	query := monitor.MainDatabase.Where("updated_at < now() - interval ? second", monitor.DelayErrorLevel)
	query.Find(&quotes)

	for _, quote := range quotes {
		errs = append(errs, errors.New(fmt.Sprintf("Symbol %s isn't updated since %s", quote.GetSymbol().Name, quote.UpdatedAt)))
	}

	return errs
}