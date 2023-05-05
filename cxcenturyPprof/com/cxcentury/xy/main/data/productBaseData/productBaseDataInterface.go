package productBaseData

import "cxcenturyPprof/com/cxcentury/xy/main/data/productBaseData/queryProductBaseDataImpl"

type ProductBaseData interface {
	QueryProductBasePresstStatusData()
	QueryProductBaseListData()
}
type DataValue struct {
	QueryProductBaseDv queryProductBaseDataImpl.DataValue
	QueryValue         *string
}
