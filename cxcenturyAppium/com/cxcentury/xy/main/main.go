package main

import (
	"fmt"
	"github.com/nichotined/appium-go-client/webdriver"
	"github.com/tebeka/selenium"
)

func main() {
	fmt.Println("nihao")
	a := selenium.Service{}
	fmt.Println(a)
	url := "http://192.168.10.101:8182"
	data := map[string]interface{}{}
	fmt.Println(webdriver.Create(url, data))

}
