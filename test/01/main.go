package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Params struct {
	Host   string
	Client int
	Num    int
	Body   string
}

func main() {
	var params Params
	flag.StringVar(&params.Host, "h", "hello", "")
	flag.IntVar(&params.Client, "c", 10, "")
	flag.IntVar(&params.Num, "n", 10, "")
	flag.StringVar(&params.Body, "d", "1233", "")
	flag.Parse()

	requestChan := make(chan interface{}, params.Num)
	writeChan := make(chan interface{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for i := 0; i < params.Num; i++ {
			requestChan <- struct {
			}{}
		}
		writeChan <- struct {
		}{}
		close(requestChan)
	}()

	var wg sync.WaitGroup
	wg.Add(params.Client)
	for i := 0; i < params.Client; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case _, ok := <-requestChan:
					if ok {
						doRequest(params)
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	<-writeChan

	if len(requestChan) <= 0 {
		cancel()
	}

	wg.Wait()
}

func doRequest(params Params) {
	fmt.Println(123)
	req, err := http.NewRequest("POST", params.Host, bytes.NewBuffer([]byte(params.Body)))
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	nowtime := time.Now()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			}
		}
	}()

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("耗时: ", time.Since(nowtime))
	statusCode := resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), statusCode)
}
