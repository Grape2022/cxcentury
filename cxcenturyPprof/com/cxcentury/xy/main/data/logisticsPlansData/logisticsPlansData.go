package logisticsPlansData

import (
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsPlansData/logisticsPlansDataImpl"
)

func AddLogisticsPlansData(
	logisticsChannelForGroupBodyList map[string][]string,
	uniqueProductCList map[string]map[string]string,
	productDataList map[string]map[string][]string,
) string {
	return logisticsPlansDataImpl.AddLogisticsPlansDataImpl(logisticsChannelForGroupBodyList, uniqueProductCList, productDataList)
}
func QueryLogisticsPlansData(addLogisticsPlansMassage map[string]string) map[string]string {
	return logisticsPlansDataImpl.QueryLogisticsPlansDataImpl(addLogisticsPlansMassage)
}
func (dv DataValue) QueryLogisticsPlansDataLogisticsMemo(dataList **[]string) ***string {
	/*dv.logisticsPlansDataImpl.QueryDataValue.Value.SearchValue,
	dv.logisticsPlansDataImpl.QueryDataValue.Value.SearchColumn,
	dv.logisticsPlansDataImpl.QueryDataValue.Value.TempSearchValue,
	dv.logisticsPlansDataImpl.QueryDataValue.Value.TempSelectValue,*/

	dv.LogisticsPlansData.QueryDataValue.Value.SearchValue = (**dataList)[0]
	dv.LogisticsPlansData.QueryDataValue.Value.SearchColumn = (**dataList)[1]
	dv.LogisticsPlansData.QueryDataValue.Value.TempSearchValue = (**dataList)[2]
	dv.LogisticsPlansData.QueryDataValue.Value.TempSelectValue = (**dataList)[3]
	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}
func (dv DataValue) QueryLogisticsPlansDataShipmentName(shipmentName **string) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.ShipmentName = **shipmentName
	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}

func (dv DataValue) QueryLogisticsPlansDataStartTimeAndEndTime(dataList **[]string) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.StartTime = (**dataList)[0]
	dv.LogisticsPlansData.QueryDataValue.Value.EndTime = (**dataList)[1]
	dv.LogisticsPlansData.QueryDataValue.Value.DateSearchType = (**dataList)[2]
	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}
func (dv DataValue) QueryLogisticsPlansDataTaxRefundType(taxRefundType **string) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.TaxRefundType = **taxRefundType
	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}

func (dv DataValue) QueryLogisticsPlansDataBatchSearchColumn(transportType **string) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.TransportType = **transportType
	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}

func (dv DataValue) QueryLogisticsPlansDataDestinationCountry(destinationCountry **string) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.DestinationCountry = **destinationCountry
	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}
func (dv DataValue) QueryLogisticsPlansDataSendCity(dataList **[]string) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.SendCity = (**dataList)[0]
	dv.LogisticsPlansData.QueryDataValue.Value.SendCityFirst = (**dataList)[1]

	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}

func (dv DataValue) QueryLogisticsPlansDataIsDelete(isDelete **int64) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.IsDelete = **isDelete
	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}
func (dv DataValue) QueryLogisticsPlansDataStatus(status **string) ***string {
	dv.LogisticsPlansData.QueryDataValue.Value.Status = **status

	dv.QueryDataValue.Data = dv.LogisticsPlansData.QueryLogisticsPlansApiDataImpl()
	return &dv.QueryDataValue.Data
}
