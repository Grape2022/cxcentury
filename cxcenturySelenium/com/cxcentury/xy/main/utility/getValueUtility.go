package utility

import (
	"fmt"
	"github.com/ser163/WordBot/generate"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

/**
取值函数工具类
*/

/*
取随机值 64
*/
func RandIntUtility(number int64) int64 {
	value := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(number)
	return value
}

/*
取随机值 64 传地址
*/
func RandIntUtilityToPath(number *int) *int64 {
	value := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(*number))
	return &value
}

/*
取随机值 float64
*/
func RandFloat64tility(number float64) float64 {
	value := rand.New(rand.NewSource(time.Now().UnixNano())).Float64() * number
	return value
}

/*
取随机值 string格式
*/
func StrcoveItoaUtilityRandIntUtility(number int64, length int) string {
	value := StrcoveFormatintUtility(RandIntUtility(number), length)
	return value
}

/*
取英文字母
生成随机单词 取首字母大写 		go get github.com/ser163/WordBot/generate
*/
func GenerateGenRandomWorldUtility() *string {
	iserr := `nil`
	var value string
	for iserr == `nil` {
		w, err := generate.GenRandomWorld(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(19), "title")
		if err != nil {
			break
		}
		//fmt.Println(w.Chance, w.Word)
		iserr = ""
		value = w.Word
	}
	return &value
}

/*
================================================================
int64 转 string
*/
func StrcoveFormatintUtility(num int64, length int) string {
	value := strconv.FormatInt(num, length)
	return value
}

// int 转string
func StrcoveItoaUtility(number int) string {
	value := strconv.Itoa(number)
	return value
}

/*
================================================================
string 转 int
*/
func StrcoveParseInt(str string) int64 {
	value, _ := strconv.ParseInt(str, 0, 0)
	//fmt.Println(reflect.TypeOf(value))
	return value
}

/*
string 转 float64
*/
func StrconvParseFloat(str string, Size int) float64 {
	floatvar, _ := strconv.ParseFloat(str, Size)
	return floatvar
}

/*
================================================================
Float64转string

		str := strconv.FormatFloat(f float64, fmt byte, prec, bitSize int)
		str 	将 float 转换成的字符串。
		f 	需要转换的 float64 类型的变量。
		fmt 	使用 f 表示不使用指数的形式。
		prec 	保留几位小数。
		bitSize 	如果为 32，表示是 float32 类型，如果是 64，表示是 float64 类型。
		fmt转为字符串后，字符串的类型：
	    'b' (-ddddp±ddd)：二进制指数
	    'e' (-d.dddde±dd)：10进制指数
	    'E' (-d.ddddE±dd)：10进制指数
	    'f' (-ddd.dddd)：无指数
	    'g' ：指数很大时使用'e'，其他使用 'f'
	    'G' ：指数很大时使用 'E'，其他使用'f'
	    'x' (-0xd.ddddp±ddd)：十六进制分数和二进制指数
	    'X' (-0Xd.ddddP±ddd)：十六进制分数和二进制指数
*/
func FmtSprintf64(value float64) string {
	str := fmt.Sprintf("%f", value)
	return str
}

func StrconvFormatFloat(value float64, prec int, bitSize int) string {
	str := strconv.FormatFloat(float64(value), 'f', prec, bitSize)
	return str
}

/*
取出interface数组字符串
*/
func ReflectValueOf() {
	//reflect.ValueOf(productInfo[`peProductName`]).Index(0).String()
	newMap := map[string]string{"你好": "2"}
	mappedToSlice := reflect.ValueOf(newMap).MapKeys()
	convertToSliceString := make([]string, len(mappedToSlice))
	for i, v := range mappedToSlice {
		convertToSliceString[i] = v.Interface().(string)
	}
	fmt.Println(len(convertToSliceString))
}

/*
把interface转换程string
*/
func InterfaceToString(inter interface{}) string {
	return fmt.Sprintf(`%s`, inter)
}

/*
把interface转换成float64
*/
func InterfaceToFloat64(inter interface{}) float64 {
	return inter.(float64)
}

// 遍历json存储map[string]interface
func ParseMap(TestMapInterface map[string]interface{}, dataMap map[string]interface{}) map[string]interface{} {

	i := 0
	for key, value := range TestMapInterface {
		switch value.(type) {
		case map[string]interface{}:
			//fmt.Println(key)
			ParseMap(value.(map[string]interface{}), dataMap)
		case string:
			if value != nil {
				dataMap[key] = value
				i += 1
				//fmt.Println(key, `:`, value)
			}
		case float64:
			if value != nil {
				dataMap[key] = value
				i += 1
				//fmt.Println(key, `:`, value)
			}
		default:
			if value != nil {
				//fmt.Println(reflect.TypeOf(value))
			}
		}
	}

	//fmt.Println(dataMap)
	return dataMap
}

// 遍历map[string]interface为map[string]string
func ForMapInterfaceToMapString(dataMap map[string]interface{}) map[string]string {
	dataMapString := map[string]string{}
	for key, value := range dataMap {
		switch value.(type) {
		case string:
			dataMapString[key] = InterfaceToString(dataMap[key])
		case float64:
			dataMapString[key] = StrconvFormatFloat(InterfaceToFloat64(dataMap[key]), 4, 64)
			//dataMapFloat64[key] = utility.InterfaceToFloat64(dataMap[key])
		}
	}
	return dataMapString
}

/*
获取外汇汇率
*/
func GetExRate(number float64, name string) float64 {
	var value float64
	if name == "USD" {
		value = number / 6.8848
	} else if name == "EUR" {
		value = number / 7.5133
	} else if name == "GBP" {
		value = number / 8.5608
	}
	//fmt.Println(`value`, reflect.TypeOf(fmt.Sprintf("%f", value)))
	return value
}

// 把外汇转成 string
func GetExRateString(str string, Size int, name string) string {
	value := StrconvFormatFloat(GetExRate(StrconvParseFloat(str, Size), name), 2, 64)
	return value
}

func init() {
	//ReflectValueOf()

}
