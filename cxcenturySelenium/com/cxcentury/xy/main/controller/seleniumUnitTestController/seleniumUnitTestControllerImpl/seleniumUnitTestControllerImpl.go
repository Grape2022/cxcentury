package seleniumUnitTestControllerImpl

import (
	"cxcenturySelenium/com/cxcentury/xy/main/concreteMothed/login"
	"cxcenturySelenium/com/cxcentury/xy/main/concreteMothed/logisticsPlans"
	"cxcenturySelenium/com/cxcentury/xy/main/concreteMothed/product"
	"cxcenturySelenium/com/cxcentury/xy/main/data/loginData"
	"cxcenturySelenium/com/cxcentury/xy/main/utility"
)

type DataValue struct {
	login          login.DataValue
	product        product.DataValue
	logisticsPlans logisticsPlans.DataValue

	loginData loginData.DataValue
}

var dv = DataValue{}

func SeleniumUnitTestControllerImpl() {
	//方法
	wd, service := utility.GetSeleniumFirefoxServerUtility()
	//登录
	phone, pwd := dv.loginData.LoginDataStar()
	dv.login.Login(&wd, phone, pwd)

	//产品
	//dv.product.AddProduct(&wd)

	//计划单
	dv.logisticsPlans.LogisticsPlans(&wd)

	utility.TimeAfter(20000)
	utility.CloseSeleniumFirefoxServerUtility(wd, service)
}
