package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				response, err := http.Get(url)
				result = Result{
					Error:    err,
					Response: response,
				}
				select {
				case <-done:
					return
				case results <- result:

				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"http://www.baidu.com", "sina.com", "http://www.aaaaaa.com", "b", "c", "d"}
	results := checkStatus(done, urls...)
	errCount := 0
	for result := range results {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("response: %v\n", result.Response)
	}
}
