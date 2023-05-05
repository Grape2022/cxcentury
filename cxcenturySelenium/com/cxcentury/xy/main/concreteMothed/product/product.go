package product

import (
	"cxcenturySelenium/com/cxcentury/xy/main/concreteMothed/product/productImpl"
	"github.com/tebeka/selenium"
)

type DataValue struct {
	ProductImplDv productImpl.DataValue
}

func (dv DataValue) AddProduct(wd *selenium.WebDriver) {

	dv.ProductImplDv.AddProductImpl(wd)

}
