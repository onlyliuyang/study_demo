package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	filename := "/Users/admin/Desktop/user_id.csv"
	ReadString(filename)
}

func ReadString(filename string)  {
	var counter map[int64]int
	counter = make(map[int64]int)

	f, _ := os.Open(filename)
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		//userId, err := strconv.Atoi(line)
		userId, err := strconv.ParseInt(line, 10, 64)
		hash := userId % 32

		if _, ok := counter[hash]; ok {
			counter[hash]++
		} else {
			counter[hash] = 1
		}
	}
	//fmt.Print(counter)

	for key, val := range counter {
		fmt.Printf("partion %d, lines :%d\n", key, val)
	}
}