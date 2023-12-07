package main

import (
	"fmt"
	"github.com/gofrs/uuid"
)

var uu = uuid.Must(uuid.NewV4())

func main() {
	for i := 0; i < 10; i++ {
		uu = uuid.Must(uuid.NewV4())
		fmt.Println(uu.String())
	}
}
