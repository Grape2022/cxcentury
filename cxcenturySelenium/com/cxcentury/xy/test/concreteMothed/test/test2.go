package test

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
	"os"
	"time"
)

func Test2() {

	// Start a WebDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}

	service, err := selenium.NewGeckoDriverService(firefoxDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	//延迟关闭服务
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}

	caps.AddFirefox(firefox.Capabilities{Args: []string{
		"window-size=1920x1080",
		//"--no-sandbox",
		//"--disable-dev-shm-usage",
		//"disable-gpu",
		//"--headless", // comment out this line to see the browser
	}})

	// Navigate to the simple playground interface.
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d", port))
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	//延迟退出firefox
	defer wd.Quit()
	// Navigate to the simple playground interface.
	if err := wd.Get("http://192.168.10.101:8182"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, ".login-tab > div:nth-child(2)")
	if err != nil {

		panic(err)
	}
	elem.Click()
	ipurtName, err := wd.FindElement(selenium.ByCSSSelector, "div.el-form-item:nth-child(1) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)")
	if err != nil {

		panic(err)
	}
	ipurtName.SendKeys("18565078921")

	password, err := wd.FindElement(selenium.ByCSSSelector, "div.el-form-item:nth-child(2) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)")
	if err != nil {

		panic(err)
	}
	password.SendKeys("YuMei000804")

	isLogin, err := wd.FindElement(selenium.ByCSSSelector, ".el-button--primary")
	if err != nil {

		panic(err)
	}
	isLogin.Click()

	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		fmt.Println(err)
		Time()
		panic(err)
	}
}
func Time() {
	<-time.After(2000000 * time.Second) //这个方法表示的是在此等待2秒，并返回一个通道，很常用
	fmt.Println("关闭")

}
func handles() {
	//查看当前窗口的handle值
	/*handle, err := wd.CurrentWindowHandle()
	if err != nil {
		panic(err)
	}
	fmt.Println(handle)
	fmt.Println("--------------------------")*/
	//查看所有网页的handle值
	/*handles, err := wd.WindowHandles()
	if err != nil {
		panic(err)
	}
	for _, handle := range handles {
		fmt.Println(handle)
	}
	fmt.Println("--------------------------")*/
}
