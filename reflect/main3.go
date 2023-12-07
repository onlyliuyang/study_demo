package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}

	//创建cat实例
	ins := cat{
		Name: "mini",
		Type: 1,
	}

	//获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	//遍历结构体所有成员

	for i := 0; i < typeOfCat.NumField(); i++ {
		//获取每个成员结构体字段
		fieldType := typeOfCat.Field(i)
		//输出成员名和tag
		fmt.Printf("name: %v, tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}

	//通过字段名找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
