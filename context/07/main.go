package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()

	valCtx := context.WithValue(rootCtx, "asong", "test01")

	go func() {
		_ = context.WithValue(valCtx, "asong", "test02")
	}()

	go func() {
		_ = context.WithValue(valCtx, "asong", "test03")
	}()

	go func() {
		for {
			fmt.Println(valCtx.Value("asong"))
		}
	}()

	go func() {
		for {
			fmt.Println(valCtx.Value("asong"))
		}
	}()
	
	time.Sleep(10 * time.Second)
}