package logisticsPlansControllerImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
	"fmt"
	"github.com/tidwall/gjson"
)

func (dv DataValue) QueryLogisticsPlansController() {
	//token传递
	dv.LogisticsPlansAPI.LoginInfo.Token = &dv.Token
	/*
		获取计划单列表
	*/
	//查询全部计划单
	dv.logisticsPlansDataImpl.QueryDataValue.Value.Status = ``
	dv.bodyString = *dv.showLogisticsPlansControllerStatus(&dv.logisticsPlansDataImpl.QueryDataValue.Value.Status)
	dv.bodyValueMapList = []bodyValueMap{}
	//fmt.Println(***dv.bodyString)

	//查询计划中计划单
	dv.logisticsPlansDataImpl.QueryDataValue.Value.Status = `1`
	dv.showLogisticsPlansControllerStatus(&dv.logisticsPlansDataImpl.QueryDataValue.Value.Status)

	//查询已完成计划单
	dv.logisticsPlansDataImpl.QueryDataValue.Value.Status = `2`
	dv.showLogisticsPlansControllerStatus(&dv.logisticsPlansDataImpl.QueryDataValue.Value.Status)

	//查询已作废计划单
	dv.logisticsPlansDataImpl.QueryDataValue.Value.IsDelete = 1
	dv.showLogisticsPlansControllerIsDelete(&dv.logisticsPlansDataImpl.QueryDataValue.Value.IsDelete)

	//查询计划单列表 按店铺

	//查询计划单列表 发货地
	dv.logisticsPlansDataImpl.QueryDataValue.Value.SendCity = `440300`
	dv.logisticsPlansDataImpl.QueryDataValue.Value.SendCityFirst = `440000`
	dv.dataList = []string{dv.logisticsPlansDataImpl.QueryDataValue.Value.SendCity, dv.logisticsPlansDataImpl.QueryDataValue.Value.SendCityFirst}
	dv.showLogisticsPlansControllerSendCity(&dv.dataList)

	//查询计划单列表 目的地
	dv.logisticsPlansDataImpl.QueryDataValue.Value.DestinationCountry = `233`
	dv.showLogisticsPlansControllerDestinationCountry(&dv.logisticsPlansDataImpl.QueryDataValue.Value.DestinationCountry)

	//查询计划单列表 运输方式
	dv.logisticsPlansDataImpl.QueryDataValue.Value.TransportType = `6`
	dv.showLogisticsPlansDataBatchSearchColumn(&dv.logisticsPlansDataImpl.QueryDataValue.Value.TransportType)

	//查询计划单列表 退税类型
	dv.logisticsPlansDataImpl.QueryDataValue.Value.TaxRefundType = `2`
	dv.showLogisticsPlansDataTaxRefundType(&dv.logisticsPlansDataImpl.QueryDataValue.Value.TaxRefundType)

	//查询计划单列表  预计提/发货时间
	dv.dateSearchType = []string{`1`, `2`, `3`, `4`}
	for k := range dv.dateSearchType {
		dv.logisticsPlansDataImpl.QueryDataValue.Value.StartTime = *(utility.RuneDateUtilityToPath(0))
		dv.logisticsPlansDataImpl.QueryDataValue.Value.EndTime = *(utility.RuneDateUtilityToPath(20))
		dv.logisticsPlansDataImpl.QueryDataValue.Value.DateSearchType = dv.dateSearchType[k]
		dv.dataList = []string{
			dv.logisticsPlansDataImpl.QueryDataValue.Value.StartTime,
			dv.logisticsPlansDataImpl.QueryDataValue.Value.EndTime,
			dv.logisticsPlansDataImpl.QueryDataValue.Value.DateSearchType,
		}
		dv.showLogisticsPlansDataStartTimeAndEndTime(&dv.dataList)
	}

	//包含标签	有异常
	//排除标识	有异常

	//查询计划单列表 货件名称
	//show random  shipmentName
	for v := range gjson.Get(***dv.bodyString, `rows`).Array() {
		dv.bodyValueMapList = append(dv.bodyValueMapList, bodyValueMap{shipmentName: map[string]string{
			`shipmentName`: gjson.Get(gjson.Get(***dv.bodyString, `rows`).Array()[v].String(), `shipmentName`).String(),
		}})
	}
	dv.number = len(dv.bodyValueMapList)
	dv.logisticsPlansDataImpl.QueryDataValue.Value.ShipmentName = dv.bodyValueMapList[*utility.RandIntUtilityToPath(&dv.number)].shipmentName[`shipmentName`]
	dv.showLogisticsPlansDataShipmentName(&dv.logisticsPlansDataImpl.QueryDataValue.Value.ShipmentName)
	//fmt.Println(`show shipmentName:`, ****dv.showLogisticsPlansDataShipmentName(&dv.logisticsPlansDataImpl.QueryDataValue.Value.ShipmentName))

	//查询计划单列表 按备注
	for v := range gjson.Get(***dv.bodyString, `rows`).Array() {
		if gjson.Get(gjson.Get(***dv.bodyString, `rows`).Array()[v].String(), `logisticsMemo`).String() != `` {
			dv.bodyValueMapList = append(dv.bodyValueMapList, bodyValueMap{shipmentName: map[string]string{
				`logisticsMemo`: gjson.Get(gjson.Get(***dv.bodyString, `rows`).Array()[v].String(), `logisticsMemo`).String(),
			}})
		}
	}
	fmt.Println(dv.bodyValueMapList)
	dv.number = len(dv.bodyValueMapList)
	dv.logisticsPlansDataImpl.QueryDataValue.Value.SearchValue = dv.bodyValueMapList[*utility.RandIntUtilityToPath(&dv.number)].shipmentName[`logisticsMemo`]
	dv.logisticsPlansDataImpl.QueryDataValue.Value.TempSearchValue = dv.logisticsPlansDataImpl.QueryDataValue.Value.SearchValue
	dv.logisticsPlansDataImpl.QueryDataValue.Value.TempSelectValue = `logisticsMemo`
	dv.dataList = []string{
		dv.logisticsPlansDataImpl.QueryDataValue.Value.SearchValue,
		dv.logisticsPlansDataImpl.QueryDataValue.Value.SearchColumn,
		dv.logisticsPlansDataImpl.QueryDataValue.Value.TempSearchValue,
		dv.logisticsPlansDataImpl.QueryDataValue.Value.TempSelectValue,
	}

	//fmt.Println(****dv.showLogisticsPlansDataLogisticsMemo(&dv.dataList))
}
func (dv DataValue) showLogisticsPlansDataLogisticsMemo(dataList *[]string) ****string {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataLogisticsMemo(&dataList)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data
	//调用请求接口
	dv.bodyString = dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
	return &dv.bodyString
}
func (dv DataValue) showLogisticsPlansDataShipmentName(shipmentName *string) ****string {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataShipmentName(&shipmentName)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data
	//调用请求接口
	dv.bodyString = dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
	return &dv.bodyString
}
func (dv DataValue) showLogisticsPlansDataStartTimeAndEndTime(dataList *[]string) {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataStartTimeAndEndTime(&dataList)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data

	//调用请求接口
	dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
}
func (dv DataValue) showLogisticsPlansDataTaxRefundType(taxRefundType *string) {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataTaxRefundType(&taxRefundType)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data

	//调用请求接口
	dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
}

func (dv DataValue) showLogisticsPlansDataBatchSearchColumn(transportType *string) {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataBatchSearchColumn(&transportType)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data

	//调用请求接口
	dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
}
func (dv DataValue) showLogisticsPlansControllerDestinationCountry(destinationCountry *string) {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataDestinationCountry(&destinationCountry)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data

	//调用请求接口
	dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
}
func (dv DataValue) showLogisticsPlansControllerSendCity(dataList *[]string) {
	//获取data并传递

	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataSendCity(&dataList)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data

	//调用请求接口
	dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
}
func (dv DataValue) showLogisticsPlansControllerIsDelete(isDelete *int64) {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataIsDelete(&isDelete)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data

	//调用请求接口
	dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
}
func (dv DataValue) showLogisticsPlansControllerStatus(status *string) ****string {
	//获取data并传递
	dv.Data = dv.LogisticsPlansData.QueryLogisticsPlansDataStatus(&status)
	dv.LogisticsPlansAPI.LoginInfo.Data = &dv.Data

	//调用请求接口
	dv.bodyString = dv.LogisticsPlansAPI.QueryLogisticsPlansAPI()
	return &dv.bodyString
}
