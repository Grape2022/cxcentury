package seleniumUnitTestController

import "cxcenturySelenium/com/cxcentury/xy/main/controller/seleniumUnitTestController/seleniumUnitTestControllerImpl"

type DataValue struct {
	DataList *map[string]string
}

func SeleniumUnitTestController() {
	seleniumUnitTestControllerImpl.SeleniumUnitTestControllerImpl()
}
