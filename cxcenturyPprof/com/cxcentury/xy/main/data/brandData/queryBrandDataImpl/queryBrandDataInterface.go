package queryBrandDataImpl

type QueryBrandData interface {
	queryBrandListDataImpl()
}
type DataValue struct {
	QueryValue     string
	QueryValuePath *string
}
