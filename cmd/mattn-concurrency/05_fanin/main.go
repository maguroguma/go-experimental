package main

import (
	"fmt"
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	uDummy "github.com/maguroguma/go-experimental/internal/usecase/dummy"
)

func main() {
	// ch := fanIn(generator("Hello"), generator("Bye"))
	ch := fanIn(another_generator("Hello"), another_generator("Bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

// goroutine を2つ作って generator を2つ束ねたりする
// new chan への書き込み順序はランダムになる
func fanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			new_ch <- <-ch1
		}
	}()
	go func() {
		for {
			new_ch <- <-ch2
		}
	}()
	return new_ch
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
