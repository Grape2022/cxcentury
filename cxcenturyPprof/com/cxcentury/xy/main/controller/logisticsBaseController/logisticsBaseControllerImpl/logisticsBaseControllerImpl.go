package logisticsBaseControllerImpl

import "cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsBases"

var dv = DataValue{}
var logisticsBasesDv = logisticsBases.DataValue{}

func (dv DataValue) QueryLogisticsBaseControllerImpl() {
	logisticsBasesDv.QueryLogisticsBase()
}
