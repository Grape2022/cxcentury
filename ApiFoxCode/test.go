package main

import (
	"fmt"
)

type DataValue struct {
	massage string
}

var dv = DataValue{}

func main() {
	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)
	dv.massage = "调用成功"
	fmt.Println(dv)

	test(a, b)

}

func test(a, b int) int {
	fmt.Println(a + b)
	return a + b
}
