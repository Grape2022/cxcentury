package productBaseControllerImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/brand"
)

func (dv DataValue) QueryApiBrandControllerImpl() *string {
	queryBrandDv := brand.DataValue{}
	queryBrandDv.Token = dv.Token

	return queryBrandDv.QueryBrandList()

}
