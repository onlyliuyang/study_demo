package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		hearbeat := make(chan interface{})
		results := make(chan time.Time)
		go func() {
			defer close(hearbeat)
			defer close(results)

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(1 * pulseInterval)

			sendPulse := func() {
				select {
				case hearbeat <- struct{}{}:
				default:

				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return hearbeat, results
	}

	done := make(chan interface{})
	time.AfterFunc(1000*time.Second, func() {
		close(done)
	})

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan)

	const timeout = 1000 * time.Second
	heartbeat, results := doWork(done, time.Second)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pluse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(timeout):
			return
		case sig := <-sigChan:
			close(done)
			fmt.Println("接受人工操作: ", sig)
		}
	}
}
