package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := requestGet()
			fmt.Println(res)
		}()
	}
	wg.Wait()
}

func requestGet() string {
	url := "http://mapi.app//applet/evaluation/list?template_id=1&platform=7&product=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err.Error()
	}
	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6ImRvdXl1eGluZ2NoZW4ifQ.eyJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3QiLCJhdWQiOiJodHRwOlwvXC9sb2NhbGhvc3QiLCJqdGkiOiJkb3V5dXhpbmdjaGVuIiwiaWF0IjoxNjk2ODE5ODA4LCJleHAiOjE3MjgzNTU4MDgsInVpZCI6MTQ2NSwicHJvZHVjdCI6MSwicGxhdGZvcm0iOjF9.kFLi2xyNeVZuTRdtzfuGTh_Pmy-1GAE_F2Ue2sk2hWs")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "mapi.app")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
