package productBaseData

import (
	"cxcenturyPprof/com/cxcentury/xy/main/data/productBaseData/addProducttBaseDataImpl"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productBaseData/queryProductBaseDataImpl"
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
	"fmt"
	"github.com/tidwall/gjson"
)

func AddProductBaseData(productDataList map[string]map[string][]string, uniqueProductCList map[string]map[string]string) string {
	return addProducttBaseDataImpl.AddProductBaseDataImpl(productDataList, uniqueProductCList)
}

var ProductBaseDataDv queryProductBaseDataImpl.DataValue

// Query Product List
func (dv DataValue) QueryProductBaseListDataPressOrderByType() queryProductBaseDataImpl.DataValue {
	//fmt.Println(dv.QueryProductBaseDv.Status == ``)
	return ProductBaseDataDv.QueryProductBaseListData()
}

// 产品状态0停售	1清仓	2在售	3开发中
func (dv DataValue) QueryProducBaseDataPresstStatus() queryProductBaseDataImpl.DataValue {
	ProductBaseDataDv.Status = dv.QueryProductBaseDv.Status
	ProductBaseDataDv.SearchColumn = `productName`
	return ProductBaseDataDv.QueryProducBaseDataPresstStatus()
}

// Query Product Press Brand
func (dv DataValue) QueryProductBaseDataPressBrand() *queryProductBaseDataImpl.DataValue {
	//Get Random BrandId And BrandName
	lens := len(gjson.Get(*dv.QueryProductBaseDv.ProductData, `rows`).Array())
	nums := utility.RandIntUtilityToPath(&lens)
	ProductBaseDataDv.BrandId = fmt.Sprintf(gjson.Get(gjson.Get(*dv.QueryProductBaseDv.ProductData, `rows`).Array()[*nums].String(), `brandId`).String())
	ProductBaseDataDv.BrandName = gjson.Get(gjson.Get(*dv.QueryProductBaseDv.ProductData, `rows`).Array()[*nums].String(), `brandName`).String()
	ProductBaseDataDv.SearchColumn = `productName`
	ProductBaseDataDv.SearchValue = ``
	return ProductBaseDataDv.QueryProductBaseListDataPath()
}

/*
Query Product Model, Press :
Product Type :model
Product SKU
Product Name
*/
func (dv DataValue) QueryProductBase() *queryProductBaseDataImpl.DataValue {
	//Get Random Product Type Model Value
	lens := len(gjson.Get(*dv.QueryProductBaseDv.ProductData, `rows`).Array())
	nums := utility.RandIntUtilityToPath(&lens)
	ProductBaseDataDv.SearchValue = gjson.Get(gjson.Get(*dv.QueryProductBaseDv.ProductData, `rows`).Array()[*nums].String(), *dv.QueryValue).String()
	//fmt.Println(`ProductBaseDataDv.SearchValue:`, ProductBaseDataDv.SearchValue)
	ProductBaseDataDv.SearchColumn = *dv.QueryValue
	return ProductBaseDataDv.QueryProductBaseListDataPath()
}
