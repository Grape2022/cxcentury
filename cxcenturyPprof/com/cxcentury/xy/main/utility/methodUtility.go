package utility

import (
	"cxcenturyPprof/com/cxcentury/xy/main/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 接口type方法
// POST,GET,PUT接口实现工具类
func (dv DataValue) MethodUtility() (body []byte) {

	//Get,Post,Put方法
	if dv.Method == `POST` {
		dv.url = config.TestUrlPath(dv.Path)[0][dv.Path]

		dv.payload = strings.NewReader(dv.Data)
		dv.req, dv.err = http.NewRequest(dv.Method, dv.url, dv.payload)

	} else if dv.Method == "GET" {
		dv.url = config.TestUrlPath(dv.Path)[0][dv.Path] + dv.Data
		dv.req, dv.err = http.NewRequest(dv.Method, dv.url, nil)
	}

	fmt.Println(dv.Name, dv.Method, dv.url)

	client := &http.Client{}

	if dv.err != nil {
		fmt.Println(dv.err)
		return
	}
	dv.req.Header.Add("Authorization", dv.Token)
	dv.req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	dv.req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(dv.req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// POST接口实现工具类
func MethodPostUtility(path, token, data, name string) (bodys string) {
	url := config.TestUrlPath(path)[0][path]
	fmt.Println(name, url)

	method := "POST"

	//fmt.Println(data)
	payload := strings.NewReader(data)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")

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
	bodys = string(body)
	return
}

// get接口实现工具类
func MethodGetUtility(path, ApiUrlPath, token, name string) []byte {
	url := config.TestUrlPath(path)[0][path] + ApiUrlPath

	fmt.Println(name, url)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}
