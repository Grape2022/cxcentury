package main

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

func main() {
	url := `http://192.168.10.101:8182/prod-api`

	loginPath := url + `/login`
	data := `{"username":"18565078921","password":"123456789+"}`
	loginMethod := `POST`
	methodDv.Data = &data
	methodDv.Method = &loginMethod
	methodDv.Url = &loginPath
	loginData := *methodDv.MethodUtility()
	token := gjson.Get(loginData, "token").String()
	methodDv.Token = &token
	methodDv.bodyString = *methodDv.FileUploadByVerify()
	fmt.Println(methodDv.bodyString)
}

// 文件上传
func (dv MethodDataValue) FileUploadByVerify() (bodys *string) {
	url := "http://192.168.10.101:8182/prod-api/logisticsChannelForgroups/importChannelQuotationGetSheetName"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(
		"C:\\Users\\Administrator\\AppData\\Roaming\\apifox\\ExternalPrograms\\com\\cxcentury\\file\\xlsx\\11.21欧美专线VIP价格（海杰运通） (1)改名字了.xlsx")
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base(
		"C:\\Users\\Administrator\\AppData\\Roaming\\apifox\\ExternalPrograms\\com\\cxcentury\\file\\xlsx\\11.21欧美专线VIP价格（海杰运通） (1)改名字了.xlsx"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", *dv.Token)
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
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
	bodyb := string(body)
	bodys = &bodyb
	return
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

var methodDv = MethodDataValue{}
