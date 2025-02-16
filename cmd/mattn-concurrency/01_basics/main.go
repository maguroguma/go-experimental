package main

import (
	"fmt"
	"time"
)

func main() {
	go regular_print("Hello")
	fmt.Println("Second print statement")
	time.Sleep(3 * time.Second)
	fmt.Println("Third print statement") // main が終わると、他の goroutine も終わる
}

func regular_print(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(1 * time.Second)
	}
}
