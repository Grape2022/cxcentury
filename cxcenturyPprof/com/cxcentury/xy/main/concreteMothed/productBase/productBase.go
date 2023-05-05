package productBase

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/productBase/addProductBaseImpl"
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/productBase/queryProductBaseImpl"
	"cxcenturyPprof/com/cxcentury/xy/main/utility/concreteMothedApiImplUtility"
)

/*
产品接口调用位置
逻辑调用区
*/
//添加产品

func AddProductBase(token string, data string) (bodys string) {
	//获取添加产品数据
	path := `/productbase`
	bodys = addProductBaseImpl.AddProductUseSKUImpl(path+"/addProductBase", token, data)
	return
}

/*
查询接口
*/

func (dv DataValue) QueryProductBase() *string {

	ProductBaseDv := queryProductBaseImpl.DataValue{}
	ProductBaseDv.Token = dv.Token
	//接口地址
	ProductBaseDv.Path = `/productbase/list`
	ProductBaseDv.Method = `GET`
	ProductBaseDv.Bodys = dv.ProductBaseDataValue.QueryValue
	return ProductBaseDv.QueryProductBaseImpl()

}
func (dv DataValue) QueryProductBaseSKUInApiUtility() *string {
	concreteMothedDv := concreteMothedApiImplUtility.DataValue{}
	concreteMothedDv.Token = dv.Token
	concreteMothedDv.Path = `/productbase/list`
	concreteMothedDv.Bodys = dv.ProductBaseDataValuePath.QueryValue
	concreteMothedDv.Method = `GET`
	return concreteMothedDv.MothedApiImpl()
}
