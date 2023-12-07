package main

import (
	"context"
	"fmt"
	"github.com/sourcegraph/conc/pool"
)

func ContextPool() {
	p := pool.New().WithMaxGoroutines(5).WithContext(context.Background()).WithCancelOnError()

	for i := 0; i < 4; i++ {
		i := i
		p.Go(func(ctx context.Context) error {
			if i == 2 {
				//return errors.New("I will cancel all other tasks")
			}
			<-ctx.Done()
			return nil
		})
	}

	err := p.Wait()
	fmt.Println(err)
}

func main() {
	ContextPool()
	//json.Marshal()
}
