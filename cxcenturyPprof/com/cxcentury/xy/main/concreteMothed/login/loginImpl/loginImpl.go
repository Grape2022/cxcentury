package loginImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"

	"github.com/tidwall/gjson"
)

func LoginUserStarImpl(path, _, data, name string) map[string]string {
	var bodyMap = make(map[string]string)
	//调用登录接口
	bodys := utility.MethodPostUtility(path, ``, data, name)

	// 方式三：不反序列化，只读取单个key，经常使用。适合特别复杂的json字符串，或者有多种if else结构的场景

	// 取得extra数组0位置的对象

	bodyMap["msg"] = gjson.Get(bodys, "msg").String()
	bodyMap["code"] = gjson.Get(bodys, "code").String()
	bodyMap["token"] = gjson.Get(bodys, "token").String()

	return bodyMap
}
