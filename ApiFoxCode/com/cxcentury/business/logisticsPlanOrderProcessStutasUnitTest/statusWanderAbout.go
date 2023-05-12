package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MethodUtility interface {
	MethodUtility()
}
type MethodDataValue struct {
	Token      *string
	Data       *string
	Name       *string
	Method     *string
	payload    *strings.Reader
	req        *http.Request
	err        error
	Url        *string
	bodyString string
}

var methodDv = MethodDataValue{}

/*
货代运单中心
*/
func main() {

	/*
		登录
	*/
	methodDv.Data = new(string)
	*methodDv.Data = `{"username":"18565078921","password":"123456789+"}`
	methodDv.Method = new(string)
	*methodDv.Method = `POST`
	methodDv.Url = new(string)
	*methodDv.Url = `http://192.168.10.101:8184/prod-api/login`
	loginData := *methodDv.MethodUtility()

	//fmt.Println(loginData)
	methodDv.Token = new(string)
	*methodDv.Token = gjson.Get(loginData, "token").String()

	/*
		获取待入库运单id
	*/
	*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/page?orderStatus=1&isDelete=0&pageSize=20&pageNum=1&orderByColumn=orderTime&isAsc=desc&transportType=&logisticsChannelForgroupId=&storageType=&destinationCountry=&groupId=&dateSearchType=1&startTime=&endTime=&isCancelType=`
	*methodDv.Method = "GET"
	*methodDv.Data = ``
	showDataorderStatus1 := *methodDv.MethodUtility()
	//fmt.Println(showDataorderStatus1)

	/*
		揽收
	*/
	if len(gjson.Get(showDataorderStatus1, "rows").Array()) > 0 {
		*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/storage`
		*methodDv.Method = "POST"
		*methodDv.Data = fmt.Sprintf(`{"storageType":1,"deliverCompany":"仙湖","deliverLink":"马大哈","deliverLinkNumber":"123","deliverCar":"112","comePickingTime":"","comePickingLocation":"","comeLink":"","comeLinkNumber":"","storageTime":"2023-05-11T16:00:00.000Z",
"logisticsOrdersId":%s}`,
			gjson.Get(gjson.Get(showDataorderStatus1, "rows").Array()[0].String(), "logisticsOrdersId"))
		storageData := *methodDv.MethodUtility()
		//fmt.Println(storageData)
		if storageData != "" {
		}
	}

	/*
		获取已入库运单id
	*/
	*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/page?orderStatus=2&isDelete=0&pageSize=20&pageNum=1&orderByColumn=orderTime&isAsc=desc&transportType=&logisticsChannelForgroupId=&storageType=&destinationCountry=&groupId=&dateSearchType=1&startTime=&endTime=&isCancelType=`
	*methodDv.Method = "GET"
	*methodDv.Data = ``
	showDataorderStatus2 := *methodDv.MethodUtility()

	/*
		核单
	*/
	if len(gjson.Get(showDataorderStatus2, "rows").Array()) > 0 {
		*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/checklist`
		*methodDv.Method = "POST"
		*methodDv.Data = fmt.Sprintf(`[{"logisticsOrdersShipmentId":"884",
"logisticsOrdersId":"%s","logisticsPlanOrderId":null,"shipmentId":"","sellerId":null,"logisticsOrdersPackages":[{"logisticsOrdersPackageId":"1375","logisticsOrdersShipmentId":"884","logisticsOrdersId":"576","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"20412","logisticsOrdersPackageId":"1375","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"20413","logisticsOrdersPackageId":"1375","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"20414","logisticsOrdersPackageId":"1375","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"20415","logisticsOrdersPackageId":"1375","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"20416","logisticsOrdersPackageId":"1375","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"20417","logisticsOrdersPackageId":"1375","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"20418","logisticsOrdersPackageId":"1375","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"20419","logisticsOrdersPackageId":"1375","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"20420","logisticsOrdersPackageId":"1375","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"20421","logisticsOrdersPackageId":"1375","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"20422","logisticsOrdersPackageId":"1375","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"20423","logisticsOrdersPackageId":"1375","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"20424","logisticsOrdersPackageId":"1375","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"20425","logisticsOrdersPackageId":"1375","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"20426","logisticsOrdersPackageId":"1375","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"20427","logisticsOrdersPackageId":"1375","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"20428","logisticsOrdersPackageId":"1375","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"20429","logisticsOrdersPackageId":"1375","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"20430","logisticsOrdersPackageId":"1375","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"20431","logisticsOrdersPackageId":"1375","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"20432","logisticsOrdersPackageId":"1375","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"20433","logisticsOrdersPackageId":"1375","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"20434","logisticsOrdersPackageId":"1375","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"20435","logisticsOrdersPackageId":"1375","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"20436","logisticsOrdersPackageId":"1375","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"20437","logisticsOrdersPackageId":"1375","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"20438","logisticsOrdersPackageId":"1375","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"20439","logisticsOrdersPackageId":"1375","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"20440","logisticsOrdersPackageId":"1375","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"20441","logisticsOrdersPackageId":"1375","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1681","logisticsOrdersPackageId":"1375","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1376","logisticsOrdersShipmentId":"884","logisticsOrdersId":"576","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"20442","logisticsOrdersPackageId":"1376","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"20443","logisticsOrdersPackageId":"1376","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"20444","logisticsOrdersPackageId":"1376","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"20445","logisticsOrdersPackageId":"1376","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"20446","logisticsOrdersPackageId":"1376","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"20447","logisticsOrdersPackageId":"1376","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"20448","logisticsOrdersPackageId":"1376","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"20449","logisticsOrdersPackageId":"1376","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"20450","logisticsOrdersPackageId":"1376","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"20451","logisticsOrdersPackageId":"1376","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"20452","logisticsOrdersPackageId":"1376","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"20453","logisticsOrdersPackageId":"1376","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"20454","logisticsOrdersPackageId":"1376","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"20455","logisticsOrdersPackageId":"1376","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"20456","logisticsOrdersPackageId":"1376","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"20457","logisticsOrdersPackageId":"1376","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"20458","logisticsOrdersPackageId":"1376","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"20459","logisticsOrdersPackageId":"1376","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"20460","logisticsOrdersPackageId":"1376","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"20461","logisticsOrdersPackageId":"1376","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"20462","logisticsOrdersPackageId":"1376","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"20463","logisticsOrdersPackageId":"1376","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"20464","logisticsOrdersPackageId":"1376","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"20465","logisticsOrdersPackageId":"1376","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"20466","logisticsOrdersPackageId":"1376","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"20467","logisticsOrdersPackageId":"1376","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"20468","logisticsOrdersPackageId":"1376","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"20469","logisticsOrdersPackageId":"1376","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"20470","logisticsOrdersPackageId":"1376","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"20471","logisticsOrdersPackageId":"1376","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1682","logisticsOrdersPackageId":"1376","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1377","logisticsOrdersShipmentId":"884","logisticsOrdersId":"576","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"20472","logisticsOrdersPackageId":"1377","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"20473","logisticsOrdersPackageId":"1377","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"20474","logisticsOrdersPackageId":"1377","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"20475","logisticsOrdersPackageId":"1377","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"20476","logisticsOrdersPackageId":"1377","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"20477","logisticsOrdersPackageId":"1377","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"20478","logisticsOrdersPackageId":"1377","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"20479","logisticsOrdersPackageId":"1377","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"20480","logisticsOrdersPackageId":"1377","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"20481","logisticsOrdersPackageId":"1377","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"20482","logisticsOrdersPackageId":"1377","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"20483","logisticsOrdersPackageId":"1377","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"20484","logisticsOrdersPackageId":"1377","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"20485","logisticsOrdersPackageId":"1377","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"20486","logisticsOrdersPackageId":"1377","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"20487","logisticsOrdersPackageId":"1377","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"20488","logisticsOrdersPackageId":"1377","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"20489","logisticsOrdersPackageId":"1377","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"20490","logisticsOrdersPackageId":"1377","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"20491","logisticsOrdersPackageId":"1377","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"20492","logisticsOrdersPackageId":"1377","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"20493","logisticsOrdersPackageId":"1377","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"20494","logisticsOrdersPackageId":"1377","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"20495","logisticsOrdersPackageId":"1377","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"20496","logisticsOrdersPackageId":"1377","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"20497","logisticsOrdersPackageId":"1377","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"20498","logisticsOrdersPackageId":"1377","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"20499","logisticsOrdersPackageId":"1377","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"20500","logisticsOrdersPackageId":"1377","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"20501","logisticsOrdersPackageId":"1377","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1683","logisticsOrdersPackageId":"1377","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]}],"logisticsOrdersPackageFixes":[{"logisticsOrdersPackageId":"1375","logisticsOrdersShipmentId":"884","logisticsOrdersId":"576","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"20412","logisticsOrdersPackageId":"1375","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"20413","logisticsOrdersPackageId":"1375","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"20414","logisticsOrdersPackageId":"1375","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"20415","logisticsOrdersPackageId":"1375","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"20416","logisticsOrdersPackageId":"1375","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"20417","logisticsOrdersPackageId":"1375","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"20418","logisticsOrdersPackageId":"1375","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"20419","logisticsOrdersPackageId":"1375","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"20420","logisticsOrdersPackageId":"1375","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"20421","logisticsOrdersPackageId":"1375","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"20422","logisticsOrdersPackageId":"1375","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"20423","logisticsOrdersPackageId":"1375","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"20424","logisticsOrdersPackageId":"1375","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"20425","logisticsOrdersPackageId":"1375","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"20426","logisticsOrdersPackageId":"1375","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"20427","logisticsOrdersPackageId":"1375","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"20428","logisticsOrdersPackageId":"1375","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"20429","logisticsOrdersPackageId":"1375","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"20430","logisticsOrdersPackageId":"1375","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"20431","logisticsOrdersPackageId":"1375","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"20432","logisticsOrdersPackageId":"1375","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"20433","logisticsOrdersPackageId":"1375","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"20434","logisticsOrdersPackageId":"1375","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"20435","logisticsOrdersPackageId":"1375","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"20436","logisticsOrdersPackageId":"1375","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"20437","logisticsOrdersPackageId":"1375","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"20438","logisticsOrdersPackageId":"1375","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"20439","logisticsOrdersPackageId":"1375","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"20440","logisticsOrdersPackageId":"1375","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"20441","logisticsOrdersPackageId":"1375","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1681","logisticsOrdersPackageId":"1375","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1376","logisticsOrdersShipmentId":"884","logisticsOrdersId":"576","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"20442","logisticsOrdersPackageId":"1376","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"20443","logisticsOrdersPackageId":"1376","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"20444","logisticsOrdersPackageId":"1376","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"20445","logisticsOrdersPackageId":"1376","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"20446","logisticsOrdersPackageId":"1376","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"20447","logisticsOrdersPackageId":"1376","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"20448","logisticsOrdersPackageId":"1376","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"20449","logisticsOrdersPackageId":"1376","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"20450","logisticsOrdersPackageId":"1376","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"20451","logisticsOrdersPackageId":"1376","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"20452","logisticsOrdersPackageId":"1376","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"20453","logisticsOrdersPackageId":"1376","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"20454","logisticsOrdersPackageId":"1376","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"20455","logisticsOrdersPackageId":"1376","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"20456","logisticsOrdersPackageId":"1376","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"20457","logisticsOrdersPackageId":"1376","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"20458","logisticsOrdersPackageId":"1376","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"20459","logisticsOrdersPackageId":"1376","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"20460","logisticsOrdersPackageId":"1376","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"20461","logisticsOrdersPackageId":"1376","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"20462","logisticsOrdersPackageId":"1376","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"20463","logisticsOrdersPackageId":"1376","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"20464","logisticsOrdersPackageId":"1376","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"20465","logisticsOrdersPackageId":"1376","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"20466","logisticsOrdersPackageId":"1376","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"20467","logisticsOrdersPackageId":"1376","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"20468","logisticsOrdersPackageId":"1376","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"20469","logisticsOrdersPackageId":"1376","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"20470","logisticsOrdersPackageId":"1376","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"20471","logisticsOrdersPackageId":"1376","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1682","logisticsOrdersPackageId":"1376","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1377","logisticsOrdersShipmentId":"884","logisticsOrdersId":"576","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"20472","logisticsOrdersPackageId":"1377","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"20473","logisticsOrdersPackageId":"1377","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"20474","logisticsOrdersPackageId":"1377","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"20475","logisticsOrdersPackageId":"1377","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"20476","logisticsOrdersPackageId":"1377","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"20477","logisticsOrdersPackageId":"1377","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"20478","logisticsOrdersPackageId":"1377","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"20479","logisticsOrdersPackageId":"1377","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"20480","logisticsOrdersPackageId":"1377","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"20481","logisticsOrdersPackageId":"1377","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"20482","logisticsOrdersPackageId":"1377","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"20483","logisticsOrdersPackageId":"1377","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"20484","logisticsOrdersPackageId":"1377","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"20485","logisticsOrdersPackageId":"1377","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"20486","logisticsOrdersPackageId":"1377","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"20487","logisticsOrdersPackageId":"1377","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"20488","logisticsOrdersPackageId":"1377","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"20489","logisticsOrdersPackageId":"1377","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"20490","logisticsOrdersPackageId":"1377","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"20491","logisticsOrdersPackageId":"1377","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"20492","logisticsOrdersPackageId":"1377","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"20493","logisticsOrdersPackageId":"1377","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"20494","logisticsOrdersPackageId":"1377","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"20495","logisticsOrdersPackageId":"1377","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"20496","logisticsOrdersPackageId":"1377","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"20497","logisticsOrdersPackageId":"1377","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"20498","logisticsOrdersPackageId":"1377","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"20499","logisticsOrdersPackageId":"1377","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"20500","logisticsOrdersPackageId":"1377","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"20501","logisticsOrdersPackageId":"1377","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1683","logisticsOrdersPackageId":"1377","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]}],"_X_ROW_KEY":"row_161"}]`,
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"))
		checklistData := *methodDv.MethodUtility()
		//fmt.Println(checklistData)
		if checklistData != "" {
		}
	}
	time.Sleep(time.Second)
	/*
		获取待出运运单id
	*/

	*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/page?orderStatus=4&isDelete=0&pageSize=20&pageNum=1&orderByColumn=orderTime&isAsc=desc&transportType=&logisticsChannelForgroupId=&storageType=&destinationCountry=&groupId=&dateSearchType=1&startTime=&endTime=&isCancelType=`
	*methodDv.Method = "GET"
	*methodDv.Data = ``
	showDataorderStatus4 := *methodDv.MethodUtility()

	/*
		直接出运
	*/
	if len(gjson.Get(showDataorderStatus4, "rows").Array()) > 0 {
		*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/directSend`
		*methodDv.Method = "POST"
		*methodDv.Data = fmt.Sprintf(`[%s]`,
			gjson.Get(gjson.Get(showDataorderStatus4, "rows").Array()[0].String(), "logisticsOrdersId"))
		directSendData := *methodDv.MethodUtility()
		//fmt.Println(directSendData)
		if directSendData != "" {
		}
	}

	/*
		获取已出运运单id
	*/
	*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/page?orderStatus=5&isDelete=0&pageSize=20&pageNum=1&orderByColumn=orderTime&isAsc=desc&transportType=&logisticsChannelForgroupId=&storageType=&destinationCountry=&groupId=&dateSearchType=1&startTime=&endTime=&isCancelType=`
	*methodDv.Method = "GET"
	*methodDv.Data = ``
	showDataorderStatus5 := *methodDv.MethodUtility()
	//fmt.Println(showDataorderStatus5)

	/*
		确认签收
	*/
	if len(gjson.Get(showDataorderStatus5, "rows").Array()) > 0 {
		*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/confirmReceipt`
		*methodDv.Method = "POST"
		*methodDv.Data = fmt.Sprintf(`[%s]`,
			gjson.Get(gjson.Get(showDataorderStatus5, "rows").Array()[0].String(), "logisticsOrdersId"))
		confirmReceiptData := *methodDv.MethodUtility()
		//fmt.Println(confirmReceiptData)
		if confirmReceiptData != "" {
		}
	} else {
		fmt.Println("货代端没有可处理运单")
	}

}

// 接口type方法
// POST,GET,PUT接口实现工具类
func (dv MethodDataValue) MethodUtility() (bodyString *string) {
	/*	fmt.Println(`dv.Method:`, *dv.Method)
		fmt.Println(`dv.Data:`, *dv.Data)
		fmt.Println(`dv.Url:`, *dv.Url)

		if nil != dv.Token {
			fmt.Println("*dv.Token", *dv.Token)
		}
	*/
	//Get,Post,Put方法
	if *dv.Method == `POST` {
		dv.payload = strings.NewReader(*dv.Data)
		dv.req, dv.err = http.NewRequest(*dv.Method, *dv.Url, dv.payload)
	} else if *dv.Method == "GET" {
		*dv.Url = *dv.Url + *dv.Data
		dv.req, dv.err = http.NewRequest(*dv.Method, *dv.Url, nil)
	}
	//fmt.Println(*dv.Method, *dv.Url)

	client := &http.Client{}

	if dv.err != nil {
		fmt.Println(dv.err)
		return
	}

	if nil != dv.Token {
		dv.req.Header.Add("Authorization", *dv.Token)
	}

	dv.req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	dv.req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(dv.req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(string(body))
	dv.bodyString = string(body)
	bodyString = &dv.bodyString
	return
}
