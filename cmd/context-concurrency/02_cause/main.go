package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())

	wg.Add(1)
	go task(ctx)
	wg.Add(1)
	go task(ctx)

	// time.Sleep(1 * time.Second)
	time.Sleep(5 * time.Second)
	cancel(errors.New("canceled by main func"))
	wg.Wait()
}

func task(ctx context.Context) {
	defer wg.Done()

	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println(context.Cause(ctx))
	}
}
