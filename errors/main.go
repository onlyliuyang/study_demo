package main

import (
	Err "errors"
	"fmt"
	"github.com/pkg/errors"
)

func t1() error {
	return Err.New("发生错误")
}

func main() {
	err0 := t1()
	err := errors.Wrap(err0, "附加消息")
	if err != nil {
		fmt.Printf("err :%+v\n", err)
	}

	fmt.Println("hello world")
}
