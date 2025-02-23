package main

import (
	"fmt"
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	uDummy "github.com/maguroguma/go-experimental/internal/usecase/dummy"
)

func main() {
	// go regular_print("Hello")
	go another_regular_print()

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

func another_regular_print() {
	container, err := uDummy.ResolvedContainer()
	if err != nil {
		panic(err)
	}

	container.Invoke(func(s sDummy.Service) {
		for i := 0; ; i++ {
			msg, _ := s.GetItem()
			fmt.Println(msg, i)
		}
	})
}
