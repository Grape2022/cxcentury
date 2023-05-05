package logisticsPlans

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsPlans/LogisticsPlansAPIImpl"
)

func AddLogisticsPlans(token, addLogisticsPlans string) map[string]string {
	path := `/logisticsPlans`
	return LogisticsPlansAPIImpl.AddLogisticsPlansImpl(token, path, addLogisticsPlans)
}
func QueryLogisticsPlans(token, logisticsPlansId string) map[string]interface{} {
	path := `/logisticsPlans`
	return LogisticsPlansAPIImpl.QueryLogisticsPlansImpl(path, `/`+logisticsPlansId, token)
}

// 查询计划单列表
func (dv DataValue) QueryLogisticsPlansAPI() ***string {
	//定义接口path
	dv.LoginInfo.Path = `/logisticsPlans/list`
	//获取url
	dv.ConfigDv.Path = &dv.LoginInfo.Path
	dv.LoginInfo.Url = (*dv.ConfigDv.TestUrlMakePath())[dv.LoginInfo.Path]
	//传参执行方法
	dv.LogisticsPlansAPIDv.APIData.Token = &dv.LoginInfo.Token
	dv.LogisticsPlansAPIDv.APIData.Url = &dv.LoginInfo.Url
	dv.LogisticsPlansAPIDv.APIData.Data = &dv.LoginInfo.Data
	dv.LogisticsPlansAPIDv.APIData.Method = "GET"
	dv.bodyString = dv.LogisticsPlansAPIDv.QueryLogisticsPlansAPIImpl()
	return &dv.bodyString
}
