package logisticsPlans

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsPlans/LogisticsPlansAPIImpl"
	"cxcenturyPprof/com/cxcentury/xy/main/config"
)

type DataValue struct {
	LoginInfo struct {
		Token **string
		Url   string
		Path  string
		Data  ****string
	}
	bodyString          **string
	LogisticsPlansAPIDv LogisticsPlansAPIImpl.DataValue
	ConfigDv            config.DataValue
}
