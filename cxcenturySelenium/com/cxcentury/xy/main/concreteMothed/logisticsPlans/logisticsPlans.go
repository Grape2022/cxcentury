package logisticsPlans

import (
	"cxcenturySelenium/com/cxcentury/xy/main/concreteMothed/logisticsPlans/logisticsPlansImpl"
	"github.com/tebeka/selenium"
)

type DataValue struct {
	logisticsPlansImplDV logisticsPlansImpl.DataValue
}

func (dv DataValue) LogisticsPlans(wd *selenium.WebDriver) {
	dv.logisticsPlansImplDV.AddLogisticsPlansImpl(wd)
}
