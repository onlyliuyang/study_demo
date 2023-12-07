package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responseChan := make(chan *http.Response, len(urls))
		go func() {
			defer close(responseChan)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err, 123)
					continue
				}

				select {
				case responseChan <- resp:
				case <-done:
					return
				}
			}
		}()
		return responseChan
	}

	urls := []string{"http://www.baidu.com", "http.sina.com", "google.com"}
	done := make(chan interface{})
	defer close(done)

	responseChan := checkStatus(done, urls...)
	for response := range responseChan {
		fmt.Printf("status of %v\n", response.Status)
	}
}
