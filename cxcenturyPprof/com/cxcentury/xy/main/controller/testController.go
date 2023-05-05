package controller

import (
	unt "cxcenturyPprof/com/cxcentury/xy/main/controller/unitTest"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsBasesData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsChannelForGroupsData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsCompanyForGroupData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsPlansData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productBaseData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productData/consumerGoodsProductData"
	f "fmt"
)

/*
流程方法控制器
用于数据传递流转终点站,负责按模块调用,和模块见数据流转
*/

func TestController() {

	//登录
	token := unt.Login()
	f.Println(`token:`, token)

	//产品
	//控制需要传递的产品参数

	productDataList := consumerGoodsProductData.ARCFOX()
	uniqueProductCList := productData.UniqueProductC()
	data := productBaseData.AddProductBaseData(productDataList, uniqueProductCList)
	unt.Product(token, data)

	//获取货代信息

	queryLogisticsCompanyForGroupData := logisticsCompanyForGroupData.QueryLogisticsCompanyForGroupData()
	logisticsCompanyForGroupBodyList := unt.CompanyForGroup(token, queryLogisticsCompanyForGroupData)

	//获取渠道信息
	queryLogisticsChannelForGroupListData := logisticsChannelForGroupsData.QueryLogisticsChannelForGroupListData(logisticsCompanyForGroupBodyList)
	logisticsChannelForGroupBodyList := unt.ChannelForGroup(token, queryLogisticsChannelForGroupListData)

	//计划单
	addLogisticsPlansData := logisticsPlansData.AddLogisticsPlansData(logisticsChannelForGroupBodyList, uniqueProductCList, productDataList)
	dataMap := unt.LogisticsPlans(token, addLogisticsPlansData)

	//物流单
	addLogistcsBasesData := logisticsBasesData.AddLogisticsBasesData(dataMap)
	unt.LogisticsBases(token, addLogistcsBasesData)
	//获取对账单信息

}
