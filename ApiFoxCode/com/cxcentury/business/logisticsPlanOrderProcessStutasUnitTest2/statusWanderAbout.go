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
	//fmt.Println(showDataorderStatus2)
	/*
		核单
	*/
	if len(gjson.Get(showDataorderStatus2, "rows").Array()) > 0 {
		*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsOrderssProcess/checklist`
		*methodDv.Method = "POST"
		*methodDv.Data = fmt.Sprintf(`[{"logisticsOrdersShipmentId":"898","logisticsOrdersId":"%s","logisticsPlanOrderId":null,"shipmentId":"","sellerId":null,"logisticsOrdersPackages":[{"logisticsOrdersPackageId":"1417","logisticsOrdersShipmentId":"898","logisticsOrdersId":"%s","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"21672","logisticsOrdersPackageId":"1417","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"21673","logisticsOrdersPackageId":"1417","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"21674","logisticsOrdersPackageId":"1417","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"21675","logisticsOrdersPackageId":"1417","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"21676","logisticsOrdersPackageId":"1417","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"21677","logisticsOrdersPackageId":"1417","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"21678","logisticsOrdersPackageId":"1417","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"21679","logisticsOrdersPackageId":"1417","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"21680","logisticsOrdersPackageId":"1417","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"21681","logisticsOrdersPackageId":"1417","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"21682","logisticsOrdersPackageId":"1417","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"21683","logisticsOrdersPackageId":"1417","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"21684","logisticsOrdersPackageId":"1417","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"21685","logisticsOrdersPackageId":"1417","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"21686","logisticsOrdersPackageId":"1417","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"21687","logisticsOrdersPackageId":"1417","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"21688","logisticsOrdersPackageId":"1417","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"21689","logisticsOrdersPackageId":"1417","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"21690","logisticsOrdersPackageId":"1417","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"21691","logisticsOrdersPackageId":"1417","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"21692","logisticsOrdersPackageId":"1417","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"21693","logisticsOrdersPackageId":"1417","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"21694","logisticsOrdersPackageId":"1417","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"21695","logisticsOrdersPackageId":"1417","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"21696","logisticsOrdersPackageId":"1417","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"21697","logisticsOrdersPackageId":"1417","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"21698","logisticsOrdersPackageId":"1417","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"21699","logisticsOrdersPackageId":"1417","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"21700","logisticsOrdersPackageId":"1417","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"21701","logisticsOrdersPackageId":"1417","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1723","logisticsOrdersPackageId":"1417","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1418","logisticsOrdersShipmentId":"898","logisticsOrdersId":"%s","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"21702","logisticsOrdersPackageId":"1418","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"21703","logisticsOrdersPackageId":"1418","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"21704","logisticsOrdersPackageId":"1418","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"21705","logisticsOrdersPackageId":"1418","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"21706","logisticsOrdersPackageId":"1418","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"21707","logisticsOrdersPackageId":"1418","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"21708","logisticsOrdersPackageId":"1418","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"21709","logisticsOrdersPackageId":"1418","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"21710","logisticsOrdersPackageId":"1418","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"21711","logisticsOrdersPackageId":"1418","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"21712","logisticsOrdersPackageId":"1418","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"21713","logisticsOrdersPackageId":"1418","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"21714","logisticsOrdersPackageId":"1418","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"21715","logisticsOrdersPackageId":"1418","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"21716","logisticsOrdersPackageId":"1418","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"21717","logisticsOrdersPackageId":"1418","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"21718","logisticsOrdersPackageId":"1418","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"21719","logisticsOrdersPackageId":"1418","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"21720","logisticsOrdersPackageId":"1418","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"21721","logisticsOrdersPackageId":"1418","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"21722","logisticsOrdersPackageId":"1418","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"21723","logisticsOrdersPackageId":"1418","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"21724","logisticsOrdersPackageId":"1418","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"21725","logisticsOrdersPackageId":"1418","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"21726","logisticsOrdersPackageId":"1418","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"21727","logisticsOrdersPackageId":"1418","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"21728","logisticsOrdersPackageId":"1418","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"21729","logisticsOrdersPackageId":"1418","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"21730","logisticsOrdersPackageId":"1418","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"21731","logisticsOrdersPackageId":"1418","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1724","logisticsOrdersPackageId":"1418","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1419","logisticsOrdersShipmentId":"898","logisticsOrdersId":"%s","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":30,"packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":67350,"totalVolume":"459.0816","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"21732","logisticsOrdersPackageId":"1419","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"21733","logisticsOrdersPackageId":"1419","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"21734","logisticsOrdersPackageId":"1419","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"21735","logisticsOrdersPackageId":"1419","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"21736","logisticsOrdersPackageId":"1419","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"21737","logisticsOrdersPackageId":"1419","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"21738","logisticsOrdersPackageId":"1419","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"21739","logisticsOrdersPackageId":"1419","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"21740","logisticsOrdersPackageId":"1419","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"21741","logisticsOrdersPackageId":"1419","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"21742","logisticsOrdersPackageId":"1419","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"21743","logisticsOrdersPackageId":"1419","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"21744","logisticsOrdersPackageId":"1419","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"21745","logisticsOrdersPackageId":"1419","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"21746","logisticsOrdersPackageId":"1419","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"21747","logisticsOrdersPackageId":"1419","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"21748","logisticsOrdersPackageId":"1419","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"21749","logisticsOrdersPackageId":"1419","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"21750","logisticsOrdersPackageId":"1419","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"21751","logisticsOrdersPackageId":"1419","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"21752","logisticsOrdersPackageId":"1419","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"21753","logisticsOrdersPackageId":"1419","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"21754","logisticsOrdersPackageId":"1419","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"21755","logisticsOrdersPackageId":"1419","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"21756","logisticsOrdersPackageId":"1419","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"21757","logisticsOrdersPackageId":"1419","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"21758","logisticsOrdersPackageId":"1419","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"21759","logisticsOrdersPackageId":"1419","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"21760","logisticsOrdersPackageId":"1419","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"21761","logisticsOrdersPackageId":"1419","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1725","logisticsOrdersPackageId":"1419","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]}],"logisticsOrdersPackageFixes":[{"logisticsOrdersPackageId":"1417","logisticsOrdersShipmentId":"898","logisticsOrdersId":"%s","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":"31","packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":69595,"totalVolume":"474.3843","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"21672","logisticsOrdersPackageId":"1417","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"21673","logisticsOrdersPackageId":"1417","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"21674","logisticsOrdersPackageId":"1417","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"21675","logisticsOrdersPackageId":"1417","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"21676","logisticsOrdersPackageId":"1417","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"21677","logisticsOrdersPackageId":"1417","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"21678","logisticsOrdersPackageId":"1417","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"21679","logisticsOrdersPackageId":"1417","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"21680","logisticsOrdersPackageId":"1417","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"21681","logisticsOrdersPackageId":"1417","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"21682","logisticsOrdersPackageId":"1417","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"21683","logisticsOrdersPackageId":"1417","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"21684","logisticsOrdersPackageId":"1417","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"21685","logisticsOrdersPackageId":"1417","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"21686","logisticsOrdersPackageId":"1417","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"21687","logisticsOrdersPackageId":"1417","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"21688","logisticsOrdersPackageId":"1417","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"21689","logisticsOrdersPackageId":"1417","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"21690","logisticsOrdersPackageId":"1417","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"21691","logisticsOrdersPackageId":"1417","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"21692","logisticsOrdersPackageId":"1417","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"21693","logisticsOrdersPackageId":"1417","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"21694","logisticsOrdersPackageId":"1417","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"21695","logisticsOrdersPackageId":"1417","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"21696","logisticsOrdersPackageId":"1417","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"21697","logisticsOrdersPackageId":"1417","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"21698","logisticsOrdersPackageId":"1417","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"21699","logisticsOrdersPackageId":"1417","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"21700","logisticsOrdersPackageId":"1417","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"21701","logisticsOrdersPackageId":"1417","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1723","logisticsOrdersPackageId":"1417","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1418","logisticsOrdersShipmentId":"898","logisticsOrdersId":"%s","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":"31","packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":69595,"totalVolume":"474.3843","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"21702","logisticsOrdersPackageId":"1418","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"21703","logisticsOrdersPackageId":"1418","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"21704","logisticsOrdersPackageId":"1418","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"21705","logisticsOrdersPackageId":"1418","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"21706","logisticsOrdersPackageId":"1418","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"21707","logisticsOrdersPackageId":"1418","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"21708","logisticsOrdersPackageId":"1418","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"21709","logisticsOrdersPackageId":"1418","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"21710","logisticsOrdersPackageId":"1418","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"21711","logisticsOrdersPackageId":"1418","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"21712","logisticsOrdersPackageId":"1418","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"21713","logisticsOrdersPackageId":"1418","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"21714","logisticsOrdersPackageId":"1418","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"21715","logisticsOrdersPackageId":"1418","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"21716","logisticsOrdersPackageId":"1418","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"21717","logisticsOrdersPackageId":"1418","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"21718","logisticsOrdersPackageId":"1418","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"21719","logisticsOrdersPackageId":"1418","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"21720","logisticsOrdersPackageId":"1418","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"21721","logisticsOrdersPackageId":"1418","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"21722","logisticsOrdersPackageId":"1418","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"21723","logisticsOrdersPackageId":"1418","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"21724","logisticsOrdersPackageId":"1418","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"21725","logisticsOrdersPackageId":"1418","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"21726","logisticsOrdersPackageId":"1418","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"21727","logisticsOrdersPackageId":"1418","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"21728","logisticsOrdersPackageId":"1418","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"21729","logisticsOrdersPackageId":"1418","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"21730","logisticsOrdersPackageId":"1418","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"21731","logisticsOrdersPackageId":"1418","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1724","logisticsOrdersPackageId":"1418","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]},{"logisticsOrdersPackageId":"1419","logisticsOrdersShipmentId":"898","logisticsOrdersId":"%s","packageName":null,"packageCode":null,"packageLength":"4.93","packageWidth":"1.94","packageHeight":"1.60","packageWeight":2245,"packageRealWeight":null,"packageCount":"31","packageIndex":null,"pcsInBox":62,"boxWeight":2245,"totalWeight":69595,"totalVolume":"474.3843","totalPrice":null,"logisticsOrdersPackageMarks":[{"logisticsOrdersPackageMarkId":"21732","logisticsOrdersPackageId":"1419","packageMark":"U000001","packageIndex":1},{"logisticsOrdersPackageMarkId":"21733","logisticsOrdersPackageId":"1419","packageMark":"U000002","packageIndex":2},{"logisticsOrdersPackageMarkId":"21734","logisticsOrdersPackageId":"1419","packageMark":"U000003","packageIndex":3},{"logisticsOrdersPackageMarkId":"21735","logisticsOrdersPackageId":"1419","packageMark":"U000004","packageIndex":4},{"logisticsOrdersPackageMarkId":"21736","logisticsOrdersPackageId":"1419","packageMark":"U000005","packageIndex":5},{"logisticsOrdersPackageMarkId":"21737","logisticsOrdersPackageId":"1419","packageMark":"U000006","packageIndex":6},{"logisticsOrdersPackageMarkId":"21738","logisticsOrdersPackageId":"1419","packageMark":"U000007","packageIndex":7},{"logisticsOrdersPackageMarkId":"21739","logisticsOrdersPackageId":"1419","packageMark":"U000008","packageIndex":8},{"logisticsOrdersPackageMarkId":"21740","logisticsOrdersPackageId":"1419","packageMark":"U000009","packageIndex":9},{"logisticsOrdersPackageMarkId":"21741","logisticsOrdersPackageId":"1419","packageMark":"U000010","packageIndex":10},{"logisticsOrdersPackageMarkId":"21742","logisticsOrdersPackageId":"1419","packageMark":"U000011","packageIndex":11},{"logisticsOrdersPackageMarkId":"21743","logisticsOrdersPackageId":"1419","packageMark":"U000012","packageIndex":12},{"logisticsOrdersPackageMarkId":"21744","logisticsOrdersPackageId":"1419","packageMark":"U000013","packageIndex":13},{"logisticsOrdersPackageMarkId":"21745","logisticsOrdersPackageId":"1419","packageMark":"U000014","packageIndex":14},{"logisticsOrdersPackageMarkId":"21746","logisticsOrdersPackageId":"1419","packageMark":"U000015","packageIndex":15},{"logisticsOrdersPackageMarkId":"21747","logisticsOrdersPackageId":"1419","packageMark":"U000016","packageIndex":16},{"logisticsOrdersPackageMarkId":"21748","logisticsOrdersPackageId":"1419","packageMark":"U000017","packageIndex":17},{"logisticsOrdersPackageMarkId":"21749","logisticsOrdersPackageId":"1419","packageMark":"U000018","packageIndex":18},{"logisticsOrdersPackageMarkId":"21750","logisticsOrdersPackageId":"1419","packageMark":"U000019","packageIndex":19},{"logisticsOrdersPackageMarkId":"21751","logisticsOrdersPackageId":"1419","packageMark":"U000020","packageIndex":20},{"logisticsOrdersPackageMarkId":"21752","logisticsOrdersPackageId":"1419","packageMark":"U000021","packageIndex":21},{"logisticsOrdersPackageMarkId":"21753","logisticsOrdersPackageId":"1419","packageMark":"U000022","packageIndex":22},{"logisticsOrdersPackageMarkId":"21754","logisticsOrdersPackageId":"1419","packageMark":"U000023","packageIndex":23},{"logisticsOrdersPackageMarkId":"21755","logisticsOrdersPackageId":"1419","packageMark":"U000024","packageIndex":24},{"logisticsOrdersPackageMarkId":"21756","logisticsOrdersPackageId":"1419","packageMark":"U000025","packageIndex":25},{"logisticsOrdersPackageMarkId":"21757","logisticsOrdersPackageId":"1419","packageMark":"U000026","packageIndex":26},{"logisticsOrdersPackageMarkId":"21758","logisticsOrdersPackageId":"1419","packageMark":"U000027","packageIndex":27},{"logisticsOrdersPackageMarkId":"21759","logisticsOrdersPackageId":"1419","packageMark":"U000028","packageIndex":28},{"logisticsOrdersPackageMarkId":"21760","logisticsOrdersPackageId":"1419","packageMark":"U000029","packageIndex":29},{"logisticsOrdersPackageMarkId":"21761","logisticsOrdersPackageId":"1419","packageMark":"U000030","packageIndex":30}],"logisticsOrdersPackageItems":[{"logisticsOrdersPackageItemId":"1725","logisticsOrdersPackageId":"1419","msku":"Qreqdf","sku":"RG230225Test1","pcs":62,"productId":"80936","productName":"RedGrape","packageMark":null}]}],"_X_ROW_KEY":"row_257"}]`,
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"),
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"),
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"),
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"),
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"),
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"),
			gjson.Get(gjson.Get(showDataorderStatus2, "rows").Array()[0].String(), "logisticsOrdersId"),
		)
		checklistData := *methodDv.MethodUtility()
		fmt.Println(checklistData)
		if checklistData != "" {
		}
	} else {
		fmt.Println(`{"msg":"没有可处理运单"}`)
	}
	time.Sleep(time.Second)

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
