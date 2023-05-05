package config

/*
登录IP配置位置
*/

//login

func urlPath() (urll []string) {
	//var urll []string

	http := `http://`
	https := `https://`
	host := `192.168.10.101`
	port := `:8182`
	api := `/prod-api`
	urll = []string{http, https, host, port, api}

	return urll
}

func TestUrlLoginPath() []map[string]string {

	url := urlPath()[0] + urlPath()[2] + urlPath()[3] + urlPath()[4]
	var urlpath []map[string]string

	/*
		loginPath	:=0
	*/
	urlpath = append(urlpath, map[string]string{`/login`: url + `/login`})
	return urlpath
}
func TestUrlPath(path string) []map[string]string {

	url := urlPath()[0] + urlPath()[2] + urlPath()[3] + urlPath()[4]
	var urlpath []map[string]string
	urlpath = append(urlpath, map[string]string{path: url + path})
	return urlpath
}

var dv = DataValue{}

func (dv DataValue) TestUrlMakePath() *map[string]string {
	dv.Http = `http://`
	dv.Https = `https://`
	dv.Host = `192.168.10.101`
	dv.Port = `:8182`
	dv.Api = `/prod-api`
	dv.Url = dv.Http + dv.Host + dv.Port + dv.Api
	dv.UrlPath = map[string]string{
		*dv.Path: dv.Url + *dv.Path,
	}
	return &dv.UrlPath
}
