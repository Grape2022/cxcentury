package utility

import (
	"cxcenturySelenium/com/cxcentury/xy/main/config"
	"fmt"
	"github.com/tebeka/selenium"
	"os"
)

func GetSeleniumFirefoxServerUtility() (selenium.WebDriver, *selenium.Service) {
	firefoxDriverPath, port := config.ConfigFirefoxBrowserFilePath()

	// Start a WebDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}

	service, err := selenium.NewGeckoDriverService(firefoxDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}

	/*caps.AddFirefox(firefox.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		//"--headless", // comment out this line to see the browser
	}})*/

	// Navigate to the simple playground interface.
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d", port))
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	return wd, service
}
func CloseSeleniumFirefoxServerUtility(wd selenium.WebDriver, service *selenium.Service) {
	//延迟关闭服务
	defer service.Stop()
	//延迟退出firefox
	defer wd.Quit()
}
