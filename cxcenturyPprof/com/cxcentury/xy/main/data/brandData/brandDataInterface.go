package brandData

type BrandData interface {
	queryBrandListData()
}
type DataValue struct {
	BD BrandData
}
