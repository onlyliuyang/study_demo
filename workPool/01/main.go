package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <- chan int, results chan <- int)  {
	for job := range jobs {
		fmt.Printf("worker(%d) start to do (%d)\n", id, job)
		time.Sleep(time.Second)
		results <- job * job
	}
}

func main()  {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for id :=0; id < 3; id++ {
		go worker(id, jobs, results)
	}

	for job :=1; job <= 50; job++ {
		jobs <- job
	}
	close(jobs)

	for i:=0; i<=50; i++ {
		<-results
	}
}