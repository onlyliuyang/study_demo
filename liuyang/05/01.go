package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
			return
		}
	}()

	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	switch local, err := local(ctx); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func local(ctx context.Context) (string, error) {
	if dealine, ok := ctx.Deadline(); ok {
		if dealine.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", errors.New("超时了哥们")
		}
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):

	}
	return "EN/US", nil
}

func genFarewell(ctx context.Context) (string, error) {
	switch local, err := local(ctx); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}
