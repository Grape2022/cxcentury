package logisticsBases

import "cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsBases/logisticsBasesAPIImpl"

var dv = DataValue{}

func AddLogisticsBases(token, data string) string {
	path := `/logisticsBases`
	return logisticsBasesAPIImpl.AddLogisticsBasesImpl(path, token, data)
}
func (dv DataValue) QueryLogisticsBase() {

}
