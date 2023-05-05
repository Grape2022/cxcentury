package logisticsPlansController

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/login"
	"cxcenturyPprof/com/cxcentury/xy/main/controller/logisticsPlansController/logisticsPlansControllerImpl"
)

type DataValue struct {
	Token            string
	ControllerImplDv logisticsPlansControllerImpl.DataValue
}

var dv = DataValue{}

func (dv DataValue) QueryLogisticsPlansController() {

	//获取token
	dv.Token = login.Login()[`token`]
	dv.ControllerImplDv.Token = &dv.Token
	dv.ControllerImplDv.QueryLogisticsPlansController()

}
