package logisticsPlansDataImpl

import "fmt"

// 计划单详情信息查询接口
func QueryLogisticsPlansDataImpl(addlogisticsPlansMassage map[string]string) map[string]string {
	//fmt.Println(`addlogisticsPlansMassage:`, addlogisticsPlansMassage)
	return addlogisticsPlansMassage
}

var dv = DataValue{}

func (dv DataValue) QueryLogisticsPlansApiDataImpl() **string {

	if dv.QueryDataValue.Value.PageNum == 0 {
		dv.QueryDataValue.Value.PageNum = 1
	}
	if dv.QueryDataValue.Value.PageSize == 0 {
		dv.QueryDataValue.Value.PageSize = 10
	}
	if dv.QueryDataValue.Value.CompanyType == 0 {
		dv.QueryDataValue.Value.CompanyType = 1
	}
	if dv.QueryDataValue.Value.DateSearchType == `` {
		dv.QueryDataValue.Value.DateSearchType = `1`
	}
	if dv.QueryDataValue.Value.OrderByColumn == `` {
		dv.QueryDataValue.Value.OrderByColumn = `createTime`
	}
	if dv.QueryDataValue.Value.SearchColumn == `` {
		dv.QueryDataValue.Value.SearchColumn = `salesman`
	}
	if dv.QueryDataValue.Value.TempSelectValue == `` {
		dv.QueryDataValue.Value.TempSelectValue = `salesman`
	}
	if dv.QueryDataValue.Value.BatchSearchColumn == `` {
		dv.QueryDataValue.Value.BatchSearchColumn = `shipmentIdString`
	}
	if dv.QueryDataValue.Value.IsAsc == `` {
		dv.QueryDataValue.Value.IsAsc = `desc`
	}
	dv.QueryDataValue.Data = dv.SetLogisticsPlansApiData()
	return &dv.QueryDataValue.Data
}

func (dv DataValue) SetLogisticsPlansApiData() *string {
	dv.QueryDataValue.MagData = fmt.Sprintf(
		"?pageNum=%d&pageSize=%d&isAsc=%s&orderByColumn=%s&shipmentName=%s&logisticsMemo=%s&salesman=%s&shopAccount=%s&warehouseManager=%s&shipmentId=%s&logisticsCompanyId=%s&logisticsChannelId=%s&logisticsChannelAliasId=%s&status=%s&fbaPlace=%s&isDelete=%d&sendCity=%s&destinationCountry=%s&dateSearchType=%s&startTime=%s&endTime=%s&sellerId=%s&sellerIds=%s&sendCityFirst=%s&inOrderTypeStr=%s&notInOrderTypeStr=%s&searchValue=%s&searchColumn=%s&payType=%s&msku=%s&fbaCode=%s&warehousingStatus=%s&payRequestStatus=%s&companyType=%d&hasTrack=%s&isDelay=%s&isCabinetRejection=%s&logisticsNo=%s&transportNo=%s&childLogisticsNo=%s&billMemo=%s&payMemo=%s&tempSearchValue=%s&tempSelectValue=%s&batchSearchColumn=%s&transportType=%s&taxRefundType=%s",
		dv.QueryDataValue.Value.PageNum,
		dv.QueryDataValue.Value.PageSize,
		dv.QueryDataValue.Value.IsAsc,
		dv.QueryDataValue.Value.OrderByColumn,
		dv.QueryDataValue.Value.ShipmentName,
		dv.QueryDataValue.Value.LogisticsMemo,
		dv.QueryDataValue.Value.Salesman,
		dv.QueryDataValue.Value.ShopAccount,
		dv.QueryDataValue.Value.WarehouseManager,
		dv.QueryDataValue.Value.ShipmentId,
		dv.QueryDataValue.Value.LogisticsCompanyId,
		dv.QueryDataValue.Value.LogisticsChannelId,
		dv.QueryDataValue.Value.LogisticsChannelAliasId,
		dv.QueryDataValue.Value.Status,
		dv.QueryDataValue.Value.FbaPlace,
		dv.QueryDataValue.Value.IsDelete,
		dv.QueryDataValue.Value.SendCity,
		dv.QueryDataValue.Value.DestinationCountry,
		dv.QueryDataValue.Value.DateSearchType,
		dv.QueryDataValue.Value.StartTime,
		dv.QueryDataValue.Value.EndTime,
		dv.QueryDataValue.Value.SellerId,
		dv.QueryDataValue.Value.SellerIds,
		dv.QueryDataValue.Value.SendCityFirst,
		dv.QueryDataValue.Value.InOrderTypeStr,
		dv.QueryDataValue.Value.NotInOrderTypeStr,
		dv.QueryDataValue.Value.SearchValue,
		dv.QueryDataValue.Value.SearchColumn,
		dv.QueryDataValue.Value.PayType,
		dv.QueryDataValue.Value.Msku,
		dv.QueryDataValue.Value.FbaCode,
		dv.QueryDataValue.Value.WarehousingStatus,
		dv.QueryDataValue.Value.PayRequestStatus,
		dv.QueryDataValue.Value.CompanyType,
		dv.QueryDataValue.Value.HasTrack,
		dv.QueryDataValue.Value.IsDelay,
		dv.QueryDataValue.Value.IsCabinetRejection,
		dv.QueryDataValue.Value.LogisticsNo,
		dv.QueryDataValue.Value.TransportNo,
		dv.QueryDataValue.Value.ChildLogisticsNo,
		dv.QueryDataValue.Value.BillMemo,
		dv.QueryDataValue.Value.PayMemo,
		dv.QueryDataValue.Value.TempSearchValue,
		dv.QueryDataValue.Value.TempSelectValue,
		dv.QueryDataValue.Value.BatchSearchColumn,
		dv.QueryDataValue.Value.TransportType,
		dv.QueryDataValue.Value.TaxRefundType,
	)
	return &dv.QueryDataValue.MagData
}
