package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 1024

	valueOfA := reflect.ValueOf(&a)

	valueOfA = valueOfA.Elem() //取出a地址的元素

	//修改a的值
	valueOfA.SetInt(1)

	//打印a的值
	fmt.Println(valueOfA.Int())
}
