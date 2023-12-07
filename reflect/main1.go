package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

func main() {
	//声明一个空结构体
	type cat struct {
	}

	//获取结构体实例的反射的类型对象
	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	//获取zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
