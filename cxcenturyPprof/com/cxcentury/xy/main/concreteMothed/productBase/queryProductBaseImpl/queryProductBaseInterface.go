package queryProductBaseImpl

type QueryProductBase interface {
	QueryProductBasePressStatusImpl()
	QueryProductBaseListImpl()
}
type DataValue struct {
	qpb    QueryProductBase
	Token  string
	Path   string
	Bodys  string
	Method string
}
