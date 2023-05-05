package main

import (
	"cxcenturyPprof/com/cxcentury/xy/main/controller"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/*
主方法调用
*/
func main() {
	//执行物流单下单流程控制器
	//controller.TestController()
	controller.Controller()

	////运行pprof
	//http.HandleFunc("/", sayHelloHandler) //    设置访问路由
	////
	//log.Fatal(http.ListenAndServe(":9299", nil))

}
func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的

	controller.TestController()
	controller.Controller()
	//timeSleep(20000)
}

func timeSleep(times int) {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < times; i++ {
		for j := 0; j < times; j++ {

			counter++
		}
	}
}

/*
*
 1. init(),方法是在其他方法之前执行的
 2. 根据导入包的顺序执行，这里 lib1、lib2，两者按导入包的顺序执行init(),然后在执行main的init方法
*/
func init() {
	//路由注册
	//http.HandleFunc("/debug/pprof/", Index)
	//http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	//http.HandleFunc("/debug/pprof/profile", Profile)
	//http.HandleFunc("/debug/pprof/symbol", Symbol)
	//http.HandleFunc("/debug/pprof/trace", Trace)

}

func Trace(writer http.ResponseWriter, request *http.Request) {

	controller.TestController()
}

func Symbol(writer http.ResponseWriter, request *http.Request) {
	controller.TestController()
}

func Profile(writer http.ResponseWriter, request *http.Request) {
	controller.TestController()
}

func Cmdline(writer http.ResponseWriter, request *http.Request) {
	controller.TestController()
}

func Index(writer http.ResponseWriter, request *http.Request) {
	controller.TestController()
}
