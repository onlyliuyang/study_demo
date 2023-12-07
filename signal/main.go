package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signalHandler(signal os.Signal) {
	if signal == syscall.SIGTERM {
		fmt.Println("Got kill signal")
		fmt.Println("Program will terminate now")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		fmt.Println("Got CTRL+C signal")
		fmt.Println("Closing.")
		os.Exit(0)
	} else {
		fmt.Println("Ignoring signal: ", signal)
	}
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	exitChan := make(chan int)

	go func() {
		for {
			s := <-sigChan
			signalHandler(s)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			os.Exit(200)
		default:
			fmt.Println("waiting")
		}
	}
	exitCode := <-exitChan
	os.Exit(exitCode)
}
