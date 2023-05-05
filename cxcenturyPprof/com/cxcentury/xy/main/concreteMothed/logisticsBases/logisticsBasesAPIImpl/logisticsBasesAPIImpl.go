package logisticsBasesAPIImpl

import "cxcenturyPprof/com/cxcentury/xy/main/utility"

/*
物流单
*/
func AddLogisticsBasesImpl(path, token, data string) string {

	return utility.MethodPostUtility(path, token, data, `addlogisticsBasesUrl:`)
}
func QueryLogisticsBaseImpl() {

}
