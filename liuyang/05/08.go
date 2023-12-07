package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		//rand.Seed(time.Now().UnixNano())
		randINt := rand.Intn(2)
		fmt.Println(randINt)
	}
}
