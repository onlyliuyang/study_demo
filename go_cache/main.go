package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	client := cache.New(5*time.Minute, 10*time.Minute)

	client.Set("foo", "bar", cache.DefaultExpiration)

	client.Set("baz", 42, cache.NoExpiration)

	foo, found := client.Get("foo")
	if found {
		fmt.Println(foo)
	}
}
