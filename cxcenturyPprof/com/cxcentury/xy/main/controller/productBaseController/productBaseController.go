package productBaseController

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/login"
	"cxcenturyPprof/com/cxcentury/xy/main/controller/productBaseController/productBaseControllerImpl"
)

type DataValue struct {
	Token string
}

var dv = DataValue{}

func QueryProductBaseController() {

	//Get Login token
	queryApiControllerDv := productBaseControllerImpl.DataValue{}
	queryApiControllerDv.Token = login.Login()["token"]
	/*
		获取品牌控制器
	*/
	BrandListData := queryApiControllerDv.QueryApiBrandControllerImpl()
	queryApiControllerDv.Data = BrandListData

	//获取产品控制器
	queryApiControllerDv.QueryApiProductControllerImpl()
	//获取计划单控制器

}
