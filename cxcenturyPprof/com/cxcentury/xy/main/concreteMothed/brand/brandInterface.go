package brand

import "cxcenturyPprof/com/cxcentury/xy/main/data/brandData/queryBrandDataImpl"

type QueryBrand interface {
	QueryBrandList()
}
type DataValue struct {
	Token               string
	QueryBrandDataValue queryBrandDataImpl.DataValue
}
