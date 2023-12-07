package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main()  {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.sina.com.cn",
	}

	for _, url := range urls {
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetch all URL.")
	}
}
