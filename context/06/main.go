package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Add(time.Second * 5))
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	subCtx, _ := context.WithCancel(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("parents is deadline")
	}

	select {
	case <-subCtx.Done():
		fmt.Println("subCtx is deadline")
	}
}
