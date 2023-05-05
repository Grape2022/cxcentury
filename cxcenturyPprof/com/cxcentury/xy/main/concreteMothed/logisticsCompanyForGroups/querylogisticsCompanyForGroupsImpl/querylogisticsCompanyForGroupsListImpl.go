package querylogisticsCompanyForGroupsImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
	"github.com/tidwall/gjson"
)

// 查询货代具体实现
func QueryLogisticsCompanyForGroupsListImpl(path, CompanyForGroup, token string) map[string][]string {

	var bodyMap = make(map[string][]string)
	body := utility.MethodGetUtility(path, CompanyForGroup, token, `QueryLogisticsCompanyForGroupsListUrl:`)
	bodys := string(body)
	// 方式三：不反序列化，只读取单个key，经常使用。适合特别复杂的json字符串，或者有多种if else结构的场景
	//获取data参数
	for _, value := range gjson.Get(bodys, `data`).Array() {
		bodyMap[`logisticsCompanyId`] = append(bodyMap[`logisticsCompanyId`], gjson.Get(value.String(), `logisticsCompanyId`).String())
		bodyMap[`companyName`] = append(bodyMap[`companyName`], gjson.Get(value.String(), `companyName`).String())
		bodyMap[`companyShortName`] = append(bodyMap[`companyShortName`], gjson.Get(value.String(), `companyShortName`).String())
		bodyMap[`companyTaxId`] = append(bodyMap[`companyTaxId`], gjson.Get(value.String(), `companyTaxId`).String())
		bodyMap[`companyCode`] = append(bodyMap[`companyCode`], gjson.Get(value.String(), `companyCode`).String())
		bodyMap[`companyMemo`] = append(bodyMap[`companyMemo`], gjson.Get(value.String(), `companyMemo`).String())
		bodyMap[`status`] = append(bodyMap[`status`], gjson.Get(value.String(), `status`).String())
		bodyMap[`isDelete`] = append(bodyMap[`isDelete`], gjson.Get(value.String(), `isDelete`).String())
		bodyMap[`isAuth`] = append(bodyMap[`isAuth`], gjson.Get(value.String(), `isAuth`).String())
	}
	// 取得extra数组0位置的对象
	//token1 := gjson.Get(bodys, "token").Array()[1]
	//fmt.Println("get raw value by key:", msg, code, token, token1.Get("address"))

	bodyMap["msg"] = []string{gjson.Get(bodys, "msg").String()}
	bodyMap["code"] = []string{gjson.Get(bodys, "code").String()}

	//fmt.Println(`bodyMap:`, bodyMap)
	return bodyMap
}
