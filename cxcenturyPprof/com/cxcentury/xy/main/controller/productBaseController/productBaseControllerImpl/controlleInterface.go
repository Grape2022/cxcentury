package productBaseControllerImpl

type ProductController interface {
	QueryApiProductControllerImpl()
}
type BrandController interface {
	QueryApiBrandControllerImpl()
}

type DataValue struct {
	Status map[string]string
	Token  string
	Data   *string
	//ProductData
}
