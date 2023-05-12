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
		fmt.Println(confirmReceiptData)
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
