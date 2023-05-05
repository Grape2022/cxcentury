package logisticsBaseController

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/login"
	"cxcenturyPprof/com/cxcentury/xy/main/controller/logisticsBaseController/logisticsBaseControllerImpl"
	"fmt"
)

type DataValue struct {
	Token string
}

var dv = DataValue{}
var ControllerImplDv = logisticsBaseControllerImpl.DataValue{}

func (dv DataValue) QueryLogisticsBaseController() {

	//获取token
	dv.Token = login.Login()[`token`]
	ControllerImplDv.Token = &dv.Token
	fmt.Println(dv.Token)
}
