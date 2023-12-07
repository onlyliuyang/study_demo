package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		//defer func(i int) {
		//	fmt.Println(i)
		//}(i)

		defer func() {
			fmt.Println(i)
		}()
	}
}
