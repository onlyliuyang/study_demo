package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ProductInfo struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func init() {
	extra.RegisterFuzzyDecoders()
}

func main() {
	jsonStr := "{\"name\":\"AppleWatchS8\",\"price\":\"3199.02\"}"
	data := ProductInfo{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
