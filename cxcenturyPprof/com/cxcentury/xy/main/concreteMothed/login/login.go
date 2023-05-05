package login

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/login/loginImpl"
	"cxcenturyPprof/com/cxcentury/xy/main/data/loginData"
)

/*
登录方法调用
*/
func Login() map[string]string {
	/*
		star
	*/

	userInfo := loginData.LoginDataStar()

	return loginUsetStar(userInfo["username"], userInfo["password"])
}

/*
登录接口调用位置
*/
func loginUsetStar(username, password string) map[string]string {

	path := `/login`
	data := `{"username":"` + username + `","password":"` + password + `"}`
	var bodyMap = loginImpl.LoginUserStarImpl(path, ``, data, `loginStarUrl:`)

	return bodyMap
}
