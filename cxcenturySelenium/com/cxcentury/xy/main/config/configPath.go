package config

const (
	firefoxDriverPath = "D:\\device\\geckodriver-v0.33.0-win64\\geckodriver.exe"
	port              = 9919
)

/*
登录IP配置位置
*/

//login

func urlPath() (urll []string) {
	//var urll []string

	http := `http://`
	https := `https://`
	host := `192.168.10.101`
	testPort := `:8182`
	urll = []string{http, https, host, testPort}

	return urll
}

func TestUrlLoginPath() []map[string]string {

	url := urlPath()[0] + urlPath()[2] + urlPath()[3]
	var urlpath []map[string]string

	/*
		loginPath	:=0
	*/
	urlpath = append(urlpath, map[string]string{`/login`: url + `/login`})
	return urlpath
}
func TestUrlPath(path string) []map[string]string {

	url := urlPath()[0] + urlPath()[2] + urlPath()[3]
	var urlpath []map[string]string
	urlpath = append(urlpath, map[string]string{path: url + path})
	return urlpath
}

// Firefox Browser  File Path to Config.
func ConfigFirefoxBrowserFilePath() (string, int) {
	//firefoxDriverPathConfig
	return firefoxDriverPath, port
}
