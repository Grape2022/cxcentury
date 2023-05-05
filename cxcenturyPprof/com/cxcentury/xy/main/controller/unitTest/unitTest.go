package unitTest

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/login"
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsBases"
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsChannelForGroups"
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsCompanyForGroups"
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsPlans"
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/productBase"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsPlansData"
)

/*
方法调用控制分类取
负责模块内数据流转
*/
/*
调用登录
*/
func Login() string {

	////获取登录token
	token := login.Login()["token"]

	return token
}

/*
调用产品
*/
func Product(token string, data string) string {

	//调用添加产品接口
	return productBase.AddProductBase(token, data)

}

/*
调用货代
*/
func CompanyForGroup(token, CompanyForGroupData string) map[string][]string {
	//调用查询渠道接口
	return logisticsCompanyForGroups.QuerylogisticsCompanyForGroupsList(token, CompanyForGroupData)
}

/*
调用渠道
*/
func ChannelForGroup(token, ChannelForGroupData string) map[string][]string {
	return logisticsChannelForGroups.QueryLogisticsChannelForGroupsList(token, ChannelForGroupData)
}

/*
调用计划单
*/
func LogisticsPlans(token, addLogisticsPlansData string) map[string]interface{} {
	addLogisticsPlansDataMassage := logisticsPlans.AddLogisticsPlans(token, addLogisticsPlansData)
	logisticsPlansId := logisticsPlansData.QueryLogisticsPlansData(addLogisticsPlansDataMassage)[`data`]

	return logisticsPlans.QueryLogisticsPlans(token, logisticsPlansId)
}

/*
调用物流单
*/
func LogisticsBases(token, data string) string {
	return logisticsBases.AddLogisticsBases(token, data)
}

func TEST() {

}
