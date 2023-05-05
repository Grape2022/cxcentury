package unitTest

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/login"
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsPlans"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsBasesData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsChannelForGroupsData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsCompanyForGroupData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsPlansData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productBaseData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/productData/consumerGoodsProductData"
	"cxcenturyPprof/com/cxcentury/xy/main/data/testingErrorTestData"
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

type dataType struct {
	token                            string
	loginData                        map[string]string
	logisticsCompanyForGroupBodyList map[string][]string
	logisticsChannelForGroupBodyList map[string][]string
	productDataList                  map[string]map[string][]string
	uniqueProductCList               map[string]map[string]string
	productBaseData                  string
	productData                      map[string]string
	QueryLogisticsPlans              map[string]interface{}
	QueryLogisticsPlansData          map[string]string
}

var data dataType

func init() {
	data.loginData = login.Login()
	data.token = data.loginData[`token`]
}

/*
登录断言
*/
func TestLogin(t *testing.T) {
	//断言200
	testingErrorTestData.TestIsMessageTypeMapString(t, data.loginData, `loginData`)

}

// 货代断言
func TestCompanyForGroup(t *testing.T) {
	queryLogisticsCompanyForGroupData := logisticsCompanyForGroupData.QueryLogisticsCompanyForGroupData()
	data.logisticsCompanyForGroupBodyList = CompanyForGroup(data.token, queryLogisticsCompanyForGroupData)
	//断言200
	testingErrorTestData.TestIsMessageTypeMapArgsString(t, data.logisticsCompanyForGroupBodyList, `ComPanyForGroupData`)

}

// 渠道断言
func TestChannelForGroup(t *testing.T) {
	queryLogisticsChannelForGroupListData := logisticsChannelForGroupsData.QueryLogisticsChannelForGroupListData(data.logisticsCompanyForGroupBodyList)
	data.logisticsChannelForGroupBodyList = ChannelForGroup(data.token, queryLogisticsChannelForGroupListData)
	//断言200
	testingErrorTestData.TestIsMessageTypeMapArgsString(t, data.logisticsChannelForGroupBodyList, `ChannelForGroupData`)
}

// 产品断言
func TestProduct(t *testing.T) {
	data.productDataList = consumerGoodsProductData.ARCFOX()
	data.uniqueProductCList = productData.UniqueProductC()
	data.productBaseData = productBaseData.AddProductBaseData(data.productDataList, data.uniqueProductCList)
	productData := Product(data.token, data.productBaseData)
	data.productData = map[string]string{
		`msg`:  gjson.Get(productData, `msg`).String(),
		`code`: gjson.Get(productData, `code`).String(),
	}
	//断言200
	testingErrorTestData.TestIsMessageTypeMapString(t, data.productData, `productData`)
	fmt.Println(data.productData)

}

// 计划单断言
func TestLogisticsPlans(t *testing.T) {

	addLogisticsPlansData := logisticsPlansData.AddLogisticsPlansData(data.logisticsChannelForGroupBodyList, data.uniqueProductCList, data.productDataList)

	addLogisticsPlansDataMassage := logisticsPlans.AddLogisticsPlans(data.token, addLogisticsPlansData)
	//断言200	添加计划单
	testingErrorTestData.TestIsMessageTypeMapString(t, addLogisticsPlansDataMassage, `addLogisticsPlans`)

	data.QueryLogisticsPlans = logisticsPlans.QueryLogisticsPlans(data.token, addLogisticsPlansDataMassage[`data`])
	fmt.Println(data.QueryLogisticsPlans[`code`], data.QueryLogisticsPlans[`msg`])
	data.QueryLogisticsPlansData = map[string]string{
		`code`: fmt.Sprint(data.QueryLogisticsPlans[`code`]),
		`msg`:  fmt.Sprint(data.QueryLogisticsPlans[`msg`]),
	}
	//断言200 按计划单id查询计划单详情
	testingErrorTestData.TestIsMessageTypeMapString(t, data.QueryLogisticsPlansData, `QueryLogisticsPlans`)
}
func TestLogisticsBases(t *testing.T) {
	addLogistcsBasesData := logisticsBasesData.AddLogisticsBasesData(data.QueryLogisticsPlans)
	LogisticsBasesString := LogisticsBases(data.token, addLogistcsBasesData)
	LogisticsBasesData := map[string]string{
		`msg`:  gjson.Get(LogisticsBasesString, `msg`).String(),
		`code`: gjson.Get(LogisticsBasesString, `code`).String(),
	}
	testingErrorTestData.TestIsMessageTypeMapString(t, LogisticsBasesData, `LogisticsBases`)
}

func TestTEST(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TEST()
		})
	}
}
