package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func doTwoRequestsAtOnce(ctx context.Context) error {
	eg, egCtx := errgroup.WithContext(ctx)
	var resp1, resp2 *http.Response
	f := func(loc string, respIn **http.Response) func() error {
		return func() error {
			reqCtx, cancel := context.WithTimeout(egCtx, time.Second*20)
			defer cancel()

			req, _ := http.NewRequest("Get", loc, nil)
			var err error
			*respIn, err = http.DefaultClient.Do(req.WithContext(reqCtx))
			if err == nil && (*respIn).StatusCode >= 500 {
				return errors.New("unexpected!")
			}
			return err
		}
	}

	eg.Go(f("http://www.baidu.com", &resp1))
	eg.Go(f("http://www.sina.com.cn", &resp2))
	return eg.Wait()
}

func main() {
	err := doTwoRequestsAtOnce(context.TODO())
	fmt.Println(err)
}
