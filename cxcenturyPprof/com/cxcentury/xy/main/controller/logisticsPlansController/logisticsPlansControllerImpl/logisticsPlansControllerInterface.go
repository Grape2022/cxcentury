package logisticsPlansControllerImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsPlans"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsPlansData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsPlansData/logisticsPlansDataImpl"
)

// 一级
type DataValue struct {
	Token *string
	Data  ***string

	dataList       []string
	timeList       []string
	dateSearchType []string
	number         int
	bodyString     ***string

	LogisticsPlansData     logisticsPlansData.DataValue
	LogisticsPlansAPI      logisticsPlans.DataValue
	logisticsPlansDataImpl logisticsPlansDataImpl.DataValue
	bodyValueMapList
}
type bodyValueMap struct {
	shipmentName map[string]string
}
type bodyValueMapList []bodyValueMap
