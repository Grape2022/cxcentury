package utility

import (
	"strings"
	"time"
)

/*
获取时间方法工具类
*/
//获取当前时间
func GetTimeUtility() *string {
	DateTimeStr := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	//fmt.Println(reflect.TypeOf(DateTimeStr), DateTimeStr)
	return &DateTimeStr
}

// 获取在当前时间基础上增加任意天数的日期
// 方法1
func GetDateTimeUtility(DayNumber int64) (DateTime string) {
	//split截取字段
	DateTime = strings.Split(time.Unix(time.Now().Unix(), 1000000000*60*60*24*DayNumber).String(), `+`)[0]
	return
}

// 方法2 较快
func RuneDateTimeUtility(DayNumber int64) (DateTime string) {
	//rune截取字段
	DateTime = string([]rune(time.Unix(time.Now().Unix(), 1000000000*60*60*24*DayNumber).String())[0:19])
	return
}

// 方法2 较快
func RuneDateTimeUtilityToPath(DayNumber int64) (DateTime *string) {
	//rune截取字段
	str := string([]rune(time.Unix(time.Now().Unix(), 1000000000*60*60*24*DayNumber).String())[0:19])
	DateTime = &str
	return
}
func RuneDateUtilityToPath(DayNumber int64, start, end int) (DateTime *string) {
	//rune截取字段
	str := string([]rune(time.Unix(time.Now().Unix(), 1000000000*60*60*24*DayNumber).String())[start:end])
	DateTime = &str
	return
}

// 在此等待 n 秒，并返回一个通道
func TimeAfter(number time.Duration) {
	<-time.After(number * time.Second)
}
