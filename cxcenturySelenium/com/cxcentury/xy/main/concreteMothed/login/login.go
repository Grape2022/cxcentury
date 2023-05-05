package login

import (
	"cxcenturySelenium/com/cxcentury/xy/main/concreteMothed/login/loginImpl"
	"github.com/tebeka/selenium"
)

type DataValue struct {
	loginImplDV loginImpl.DataValue
}

func (dv DataValue) Login(wd *selenium.WebDriver, phoneNumber, password string) {
	dv.loginImplDV.LoginImpl(wd, phoneNumber, password)

	//getPageSource, err := wd.PageSource()
	//fmt.Println(getPageSource)
	return
}
