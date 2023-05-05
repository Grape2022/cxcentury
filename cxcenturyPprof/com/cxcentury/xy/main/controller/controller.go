package controller

import (
	"cxcenturyPprof/com/cxcentury/xy/main/controller/logisticsBaseController"
	"cxcenturyPprof/com/cxcentury/xy/main/controller/logisticsPlansController"
	"cxcenturyPprof/com/cxcentury/xy/main/controller/productBaseController"
)

type DataValue struct {
	LogisticsBaseDv  logisticsBaseController.DataValue
	LogisticsPlansDv logisticsPlansController.DataValue
}

var dv = DataValue{}

func Controller() {
	//执行查询接口控制器
	//查询产品控制器
	productBaseController.QueryProductBaseController()
	//查询物流控制器

	dv.LogisticsPlansDv.QueryLogisticsPlansController()
	//查询物流控制器
	dv.LogisticsBaseDv.QueryLogisticsBaseController()

}
