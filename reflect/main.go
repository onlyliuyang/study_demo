package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string	`json:"user_name"`
	Age	int
}

func main()  {
	//var num int = 100

	//rVal := reflect.ValueOf(num)
	//rType := reflect.TypeOf(num)
	//fmt.Println(rVal)
	//fmt.Println(rType)
	//
	////将rVal转换成interface()
	//iv := rVal.Interface()
	////将interface通过断言转成需要的类型
	//num2 := iv.(int)
	//fmt.Println(num2)

	stu := &Student{
		Name:"tom",
		Age:20,
	}

	rType := reflect.TypeOf(stu)
	//获取指针类型的原素类型
	e := rType.Elem()
	//显示指针变量指向元素的类型名称和种类
	fmt.Printf("name: %v, kind: %v\n", e.Name(), e.Kind())
}