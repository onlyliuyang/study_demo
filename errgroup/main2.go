package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	for index :=0; index < 3; index++ {
		indexTemp := index

		//新建子协程
		group.Go(func() error {
			fmt.Printf("indexTemp = %d", indexTemp)
			//第一个协程
			if indexTemp == 0 {
				fmt.Println("indexTemp == 0 start")
				fmt.Println("indexTemp == 0 end")
			} else if indexTemp == 1 {
				fmt.Println("indexTemp == 1 start")
				//time.Sleep(1 * time.Second)
				cancel()	//第二个协程异常退出
				fmt.Println("indexTemp == 1 error")
			} else if indexTemp == 2 {
				fmt.Println("indexTemp == 2 start")
				//time.Sleep(1 * time.Second)

				err := CheckGroutineErr(errCtx)
				if err != nil {
					return err
				}
				fmt.Println("indexTemp == 2 end")
			}
			return nil
		})
	}

	err := group.Wait()
	if err == nil {
		fmt.Println("都完成了")
	} else {
		fmt.Println("get error: %v", err)
	}
}

func CheckGroutineErr(errContext context.Context) error {
	select {
	case <-errContext.Done():
		return errContext.Err()
	default:
		return nil
	}
}