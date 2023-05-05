package loginData

/*
登录信息模板
*/
func LoginDataStar() map[string]string {
	uname := "18565078921"
	pwd := "YuMei000804"
	loginList := make(map[string]string)
	loginList["username"] = uname
	loginList["password"] = pwd

	return loginList
}
