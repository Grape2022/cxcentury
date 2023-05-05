package LogisticsPlansAPIImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
	"encoding/json"
	"github.com/tidwall/gjson"
)

// 添加计划单
func AddLogisticsPlansImpl(token, path, data string) map[string]string {
	bodyMap := map[string]string{}

	bodys := utility.MethodPostUtility(path, token, data, `addlogisticsPlansUrl:`)
	bodyMap[`msg`] = gjson.Get(bodys, `msg`).String()
	bodyMap[`code`] = gjson.Get(bodys, `code`).String()
	bodyMap[`data`] = gjson.Get(bodys, `data`).String()

	return bodyMap
}
func QueryLogisticsPlansImpl(path, CompanyForGroup, token string) map[string]interface{} {

	//获取计划单信息
	dataMap := make(map[string]interface{})
	body := utility.MethodGetUtility(path, CompanyForGroup, token, `QueryLogisticsPlansUrl`)

	TestMapInterface := map[string]interface{}{}
	//转化[]byte为json格式=：map[string]interface
	err := json.Unmarshal([]byte(body), &TestMapInterface)
	if err != nil {
		panic(err)
	}
	dataMap = utility.ParseMap(TestMapInterface, map[string]interface{}{})

	return dataMap
}
func (dv DataValue) QueryLogisticsPlansAPIImpl() **string {
	dv.Method.APIMethod.Url = dv.APIData.Url
	dv.Method.APIMethod.Data = ****dv.APIData.Data
	dv.Method.APIMethod.Token = **dv.APIData.Token
	dv.Method.APIMethod.Method = &dv.APIData.Method
	dv.bodyString = dv.Method.APIMethod.MethodUtility()
	return &dv.bodyString
}
