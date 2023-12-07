package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {
	}

	ins := &cat{}
	typeOfCat := reflect.TypeOf(ins)
	fmt.Printf("name : %v kind : %v\n", typeOfCat.Name(), typeOfCat.Kind())

	//取类型的原素
	typeOfCat = typeOfCat.Elem()
	//显示反射类型对象的名称和种类
	fmt.Printf("name : %v kind : %v\n", typeOfCat.Name(), typeOfCat.Kind())
}
