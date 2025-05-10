package main

import (
	"fmt"
	"time"

	"github.com/avast/retry-go"
)

var g = 0

func main() {
	startTime := time.Now()

	err := retry.Do(
		func() error {
			elapsed := time.Since(startTime).Seconds()
			fmt.Printf("Elapsed time: %.2f seconds\n", elapsed)
			err := f()
			return err
		},
		retry.Attempts(10),
		retry.Delay(1*time.Second), // 指数増加のベースの遅延時間
		retry.DelayType(retry.CombineDelay(retry.BackOffDelay, retry.RandomDelay)),
	)

	fmt.Printf("Final error: %v\n", err)
}

func f() error {
	g++
	fmt.Printf("g = %d\n", g)

	if g >= 4 {
		return nil
	}
	return fmt.Errorf("error %d", g)
}
