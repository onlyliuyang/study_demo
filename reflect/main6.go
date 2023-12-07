package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	x := 2
	//a := reflect.ValueOf(2)
	//b := reflect.ValueOf(x)
	d := reflect.ValueOf(&x).Elem()

	//fmt.Println(a.CanAddr())
	//fmt.Println(b.CanAddr())
	//fmt.Println(c.CanAddr())
	//fmt.Println(d.CanAddr())

	fmt.Println(x)
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x)

	d.Set(reflect.ValueOf(4))
	fmt.Println(x)

	w := 2
	b := reflect.ValueOf(&w).Elem()
	b.Set(reflect.ValueOf(3))

	//json.Unmarshal()
	json.Encoder{}
}
