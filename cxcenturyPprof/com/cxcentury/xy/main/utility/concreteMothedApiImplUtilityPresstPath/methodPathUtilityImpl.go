package concreteMothedApiImplUtilityPresstPath

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 接口type方法
// POST,GET,PUT接口实现工具类
func (dv DataValue) MethodUtility() (bodyString *string) {

	//Get,Post,Put方法
	if *dv.Method == `POST` {
		dv.payload = strings.NewReader(*dv.Data)
		dv.req, dv.err = http.NewRequest(*dv.Method, *dv.Url, dv.payload)
	} else if *dv.Method == "GET" {
		*dv.Url = *dv.Url + *dv.Data
		dv.req, dv.err = http.NewRequest(*dv.Method, *dv.Url, nil)
	}
	fmt.Println(*dv.Method, *dv.Url)

	client := &http.Client{}

	if dv.err != nil {
		fmt.Println(dv.err)
		return
	}
	dv.req.Header.Add("Authorization", *dv.Token)
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
