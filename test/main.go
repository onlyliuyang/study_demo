package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
)

var (
	counter    uint32 = 0
	errCourter uint32 = 0
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			resp, err := receiveProduct()
			respMap := make(map[string]interface{})
			json.Unmarshal([]byte(resp), &respMap)
			if respMap["code"].(float64) == 20000 {
				atomic.AddUint32(&counter, 1)
			} else {
				atomic.AddUint32(&errCourter, 1)
			}
			fmt.Println(resp, err)
		}()
	}
	wg.Wait()
	fmt.Println(counter, errCourter)
}

func receiveProduct() (string, error) {
	url := "https://test-mapi.douyuxingchen.com/activity/welfare/test-receive-product"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var accessToken string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6ImRvdXl1eGluZ2NoZW4ifQ.eyJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3QiLCJhdWQiOiJodHRwOlwvXC9sb2NhbGhvc3QiLCJqdGkiOiJkb3V5dXhpbmdjaGVuIiwiaWF0IjoxNjgwMDg1MjA5LCJleHAiOjE3MTE2MjEyMDksInVpZCI6MTQ2NSwicHJvZHVjdCI6MSwicGxhdGZvcm0iOjB9.neqGY6AfH9sUues-WOSmKmGScE7dXC39QloI9BvPwNE"
	req.Header.Set("Authorization", accessToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
