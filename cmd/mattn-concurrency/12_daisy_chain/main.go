package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 10000

	var channels [n + 1]chan int
	for i := range channels {
		channels[i] = make(chan int)
	}

	for i := 0; i < n; i++ {
		go f(channels[i], channels[i+1])
	}

	// 末尾に流すことで連鎖的に goroutine が流れだす
	// go func(c chan<- int) { c <- 1 }(channels[n])
	func(c chan<- int) { c <- 5 }(channels[n]) // goroutine じゃなくても動く

	fmt.Println(<-channels[0])
}
