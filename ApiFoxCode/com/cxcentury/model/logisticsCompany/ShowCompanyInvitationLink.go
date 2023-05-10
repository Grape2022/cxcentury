package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
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
	fmt.Println(loginData)
	methodDv.Token = new(string)
	*methodDv.Token = gjson.Get(loginData, "token").String()
	fmt.Println(*methodDv.Token)
	/*
		绑定报价业务员
	*/
	*methodDv.Url = `http://192.168.10.101:8184/prod-api/logisticsCompanySalesmans/addSalesman`
	*methodDv.Data = `{"param":"E7DFDF07467FE15A018B6CBF5CBEDFBBA823B107F18FAC400A8557F51CC26A52","userId":162,"userName":"18565078921"}`
	data := *methodDv.MethodUtility()
	fmt.Println(`data:`, data)
}

// 接口type方法
// POST,GET,PUT接口实现工具类
func (dv MethodDataValue) MethodUtility() (bodyString *string) {
	//fmt.Println(`dv.Method:`, *dv.Method)
	//fmt.Println(`dv.Data:`, *dv.Data)
	//fmt.Println(`dv.Url:`, *dv.Url)
	//
	//fmt.Println("*dv.Token", dv.Token)

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
