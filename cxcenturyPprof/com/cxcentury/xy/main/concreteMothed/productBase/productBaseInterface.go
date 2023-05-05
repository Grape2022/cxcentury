package productBase

import "cxcenturyPprof/com/cxcentury/xy/main/data/productBaseData/queryProductBaseDataImpl"

type QueryPrductBase interface {
	QueryProductBasePresstStatus()
}
type DataValue struct {
	Token                    string
	ProductBaseDataValue     queryProductBaseDataImpl.DataValue
	ProductBaseDataValuePath *queryProductBaseDataImpl.DataValue
}
