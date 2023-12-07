package main

import (
	"context"
	"fmt"
	"time"
)

func testWithCancel(t int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("testWithCancel.Done: ", ctx.Err())
	case e := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("testWithCancel: ", e)
	}
	return
}

func testWithDeadline(t int) {
	ctx := context.Background()
	dl := time.Now().Add(time.Duration(1 * t) * time.Second)
	ctx, cancel := context.WithDeadline(ctx, dl)
	defer cancel()

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("testWithDeadline.Done: ", ctx.Err())
	case e := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("testWithDeadline: ", e)
	}
	return
}

func testWithTimeout(t int)  {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(t) * time.Second)
	defer cancel()

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	select {
	case <- ctx.Done():
		fmt.Println("testWithTimeout.Done: ", ctx.Err())
	case e := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("testWithTimeout: ", e)
	}
	return
}

func main() {
	var t int = 2
	//testWithCancel(t)

	//testWithDeadline(t)
	testWithTimeout(t)
}
