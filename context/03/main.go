package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET", "https://www.baidu.com/aa/bb/cc", nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
	req.WithContext(ctx)
	defer cancel()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
