package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlStr string = "https://test-m.douyuxingchen.com/douyuxingchen/bestv-exchange/exchange-center"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println(escapeUrl)
	fmt.Println()
}
