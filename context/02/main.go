package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	url     string
	timeout = 1
	wg      sync.WaitGroup
)

type information struct {
	r   *http.Response
	err error
}

func connect(ctx context.Context) error {
	defer wg.Done()

	info := make(chan information, 1)
	tr := &http.Transport{}
	httpClient := http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)

	go func() {
		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Printf("request error:%s\n", err)
			info <- information{nil, err}
			return
		} else {
			info <- information{res, nil}
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("request is cancelled!!, %s", ctx.Err())
	case ok := <-info:
		err := ok.err
		r := ok.r
		if err != nil {
			fmt.Printf("ERROR: \n", err)
			return err
		}

		defer r.Body.Close()
		realInfo, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("ERROR: \n", err)
			return err
		}

		fmt.Printf("Response: %s\n", realInfo)
	}
	return nil
}

func main()  {
	url = "https://baidu.com"
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout) * time.Second)
	defer cancel()

	fmt.Printf("connect to %s\n", url)
	wg.Add(1)
	go connect(ctx)
	wg.Wait()
	fmt.Println("End...")
}
