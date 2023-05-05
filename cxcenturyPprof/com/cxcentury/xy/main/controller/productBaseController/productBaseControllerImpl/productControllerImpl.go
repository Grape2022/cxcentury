package productBaseControllerImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/productBase"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productBaseData"
)

//var QueryProductDataDv = productBaseData.DataValue{}
//var productBaseDv = productBaseController.DataValue{}

var DV struct {
	QueryValue         []string
	QueryProductDataDv productBaseData.DataValue
	ProductBaseDv      productBase.DataValue
}

func (dv DataValue) QueryApiProductControllerImpl() {

	DV.ProductBaseDv.Token = dv.Token
	/*
		Query SKU
		productBaseController API List query
	*/
	//Product List
	DV.ProductBaseDv.ProductBaseDataValue = DV.QueryProductDataDv.QueryProductBaseListDataPressOrderByType()
	dataJson := DV.ProductBaseDv.QueryProductBase()
	DV.QueryProductDataDv.QueryProductBaseDv.ProductData = dataJson

	//Product Type QueryValue
	DV.QueryValue = []string{
		`productName`,
		`sku`,
		`model`,
		`purchaser`,
		//`developer`,
		//`developer`,
		//`duty`,
		//`creator`,
	}
	for i := 0; i < len(DV.QueryValue); i++ {
		DV.QueryProductDataDv.QueryValue = &DV.QueryValue[i]
		DV.ProductBaseDv.ProductBaseDataValuePath = DV.QueryProductDataDv.QueryProductBase()
		DV.ProductBaseDv.QueryProductBaseSKUInApiUtility()
	}

	//productBaseController Api Presst Brand Query
	DV.QueryProductDataDv.QueryProductBaseDv.ProductData = dv.Data
	DV.ProductBaseDv.ProductBaseDataValuePath = DV.QueryProductDataDv.QueryProductBaseDataPressBrand()
	DV.ProductBaseDv.QueryProductBaseSKUInApiUtility()

	//productBaseController API Presst Status Query			status in 0,1,2,3
	dv.Status = map[string]string{
		`0`: `停售`,
		`1`: `清仓`,
		`2`: `在售`,
		`3`: `开发中`,
	}
	for k, _ := range dv.Status {

		DV.QueryProductDataDv.QueryProductBaseDv.Status = k
		DV.ProductBaseDv.ProductBaseDataValue = DV.QueryProductDataDv.QueryProducBaseDataPresstStatus()
		DV.ProductBaseDv.QueryProductBase()
	}

	/*
		Query SPU
	*/

}
