package addProductBaseImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
)

/*
添加产品具体实现
*/
func AddProductUseSKUImpl(path, token, data string) (bodys string) {
	bodys = utility.MethodPostUtility(path, token, data, `addProductUrl:`)

	return
}
