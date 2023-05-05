package loginImpl

import (
	"cxcenturySelenium/com/cxcentury/xy/main/config"
	"cxcenturySelenium/com/cxcentury/xy/main/utility"
	"github.com/tebeka/selenium"
	"strings"
	"time"
)

type DataValue struct {
	DataMapList []map[string][]string
}

//type DataMap struct {
//	Name            string
//	CSSSelectorPath string
//	Value           string
//}

func (dv DataValue) LoginImpl(wd *selenium.WebDriver, phoneNumber, password string) {
	path := `/#/cxc-logistics-web/login`

	// Navigate to the simple playground interface.
	if err := (*wd).Get(config.TestUrlPath(path)[0][path]); err != nil {
		panic(err)
	}
	dv.DataMapList = make([]map[string][]string, 0)
	dv.DataMapList = []map[string][]string{
		{`PasswordLoginClickCSS`: {`.login-tab > div:nth-child(2)`}},
		{`PhoneNumberInputCSS`: {`div.el-form-item:nth-child(1) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)`,
			phoneNumber}},
		{`PasswordInputCSS`: {`div.el-form-item:nth-child(2) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)`,
			password}},
		{`LoginSubmitClickCSS`: {`.el-button--primary`}},
		{`DataPageOkIsSecondClickCSS`: {`span.center-button:nth-child(4)`}},
	}
	//dv.DataMapList = append(dv.DataMapList, struct {
	//	Name            string
	//	CSSSelectorPath string
	//	Value           string
	//}{Name: `ClickPasswordLogin`, CSSSelectorPath: `.login-tab > div:nth-child(2)`, Value: ``},
	//	struct {
	//		Name            string
	//		CSSSelectorPath string
	//		Value           string
	//	}{Name: `InputPhoneNumber`, CSSSelectorPath: `div.el-form-item:nth-child(1) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)`, Value: phoneNumber},
	//	struct {
	//		Name            string
	//		CSSSelectorPath string
	//		Value           string
	//	}{Name: `InputPassword`, CSSSelectorPath: `div.el-form-item:nth-child(2) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)`, Value: password},
	//	struct {
	//		Name            string
	//		CSSSelectorPath string
	//		Value           string
	//	}{Name: `ClickLogin`, CSSSelectorPath: `.el-button--primary`, Value: ``},
	//)
	for i := 0; i < len(dv.DataMapList); i++ {
		for k := range dv.DataMapList[i] {
			if strings.Contains(k, `IsSecond`) {
				time.Sleep(time.Second * 1)
			}
			switch true {
			case strings.Contains(k, `Click`) && strings.Contains(k, `CSS`):
				utility.ElementClickCSS(wd, &dv.DataMapList[i][k][0])
			case strings.Contains(k, `Click`) && strings.Contains(k, `XPATH`):
				utility.ElementClickXPATH(wd, &dv.DataMapList[i][k][0])
			case strings.Contains(k, `Click`) && strings.Contains(k, `TagName`):
				utility.ElementClickByTagName(wd, &dv.DataMapList[i][k][0])
			case strings.Contains(k, `Input`) && strings.Contains(k, `CSS`):
				utility.ElementSendKeysCSS(wd, &dv.DataMapList[i][k][0], &dv.DataMapList[i][k][1])
			case strings.Contains(k, `Input`) && strings.Contains(k, `XPATH`):
				utility.ElementSendKeysXPATH(wd, &dv.DataMapList[i][k][0], &dv.DataMapList[i][k][1])
			case strings.Contains(k, `Input`) && strings.Contains(k, `TagName`):
				utility.ElementSendKeysByTagName(wd, &dv.DataMapList[i][k][0], &dv.DataMapList[i][k][1])
			}
		}
	}

	return
}
