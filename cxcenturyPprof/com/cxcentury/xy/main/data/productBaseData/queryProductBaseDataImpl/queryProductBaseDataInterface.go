package queryProductBaseDataImpl

/*
产品查询数据模板
*/
type QueryProductBaseData interface {
	QueryProducBasePresstStatusData() DataValue
	QueryProductBaseListData() DataValue
}
type DataValue struct {
	QueryValue    string
	PageNum       int64
	PageSize      int64
	IsAsc         string
	OrderByColumn string
	IsCombo       string
	Status        string
	BrandId       string
	SearchColumn  string
	SearchValue   string
	JointQuery    string
	StockTypeTab  string
	BrandName     string
	CategoryName  string
	ProductName   string
	OrderByKey    string
	OrderByType   string
	ProductData   *string
}
