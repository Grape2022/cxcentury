package logisticsPlansDataImpl

type LogisticsPlansDataInterface interface {
}
type DataValue struct {
	QueryDataValue struct {
		Data    *string
		MagData string
		Value   struct {
			PageNum                 int64
			PageSize                int64
			IsAsc                   string
			OrderByColumn           string
			ShipmentName            string
			LogisticsMemo           string
			Salesman                string
			ShopAccount             string
			WarehouseManager        string
			ShipmentId              string
			LogisticsCompanyId      string
			LogisticsChannelId      string
			LogisticsChannelAliasId string
			Status                  string
			FbaPlace                string
			IsDelete                int64
			SendCity                string
			DestinationCountry      string
			DateSearchType          string
			StartTime               string
			EndTime                 string
			SellerId                string
			SellerIds               string
			SendCityFirst           string
			InOrderTypeStr          string
			NotInOrderTypeStr       string
			SearchValue             string
			SearchColumn            string
			PayType                 string
			Msku                    string
			FbaCode                 string
			WarehousingStatus       string
			PayRequestStatus        string
			CompanyType             int64
			HasTrack                string
			IsDelay                 string
			IsCabinetRejection      string
			LogisticsNo             string
			TransportNo             string
			ChildLogisticsNo        string
			BillMemo                string
			PayMemo                 string
			TempSearchValue         string
			TempSelectValue         string
			BatchSearchColumn       string
			TransportType           string //新增 运输方式
			TaxRefundType           string //新增 退税类型	1,2,3

		}
	}
}
