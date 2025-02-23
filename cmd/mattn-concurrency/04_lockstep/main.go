package main

import (
	"fmt"
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	uDummy "github.com/maguroguma/go-experimental/internal/usecase/dummy"
)

func main() {
	// ch1 := generator("Hello")
	// ch2 := generator("Bye")
	ch1 := another_generator("Hello")
	ch2 := another_generator("Bye")

	for i := 0; i < 5; i++ {
		// ここで必ず順序制御される
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func another_generator(arg string) <-chan string {
	ch := make(chan string)

	container, err := uDummy.ResolvedContainer()
	if err != nil {
		panic(err)
	}

	go container.Invoke(func(s sDummy.Service) {
		for i := 0; ; i++ {
			msg, _ := s.GetItem()
			ch <- fmt.Sprintf("%s %s %d", arg, msg, i)
		}
	})

	return ch
}
