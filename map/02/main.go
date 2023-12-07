package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	//for key, value := range scoreMap {
	//	fmt.Println(key, value)
	//}

	var keys []string = make([]string, 0, 100)
	for key, _ := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	fmt.Println(scoreMap)
}
