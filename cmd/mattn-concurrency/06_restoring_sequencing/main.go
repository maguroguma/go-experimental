package main

import (
	"fmt"
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	uDummy "github.com/maguroguma/go-experimental/internal/usecase/dummy"
)

type Message struct {
	str   string
	block chan int // メッセージオブジェクトに共用のチャンネルを持たせる
}

func main() {
	// ch := fanIn(generator("Hello"), generator("Bye"))
	ch := fanIn(another_generator("Hello"), another_generator("Bye"))
	for i := 0; i < 5; i++ {
		msg1 := <-ch
		fmt.Println(msg1.str)

		msg2 := <-ch
		fmt.Println(msg2.str)

		<-msg1.block
		<-msg2.block
	}
}

func fanIn(ch1, ch2 <-chan Message) <-chan Message {
	new_ch := make(chan Message)
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

func generator(msg string) <-chan Message {
	ch := make(chan Message)
	blockingStep := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("%s %d", msg, i), blockingStep}
			time.Sleep(time.Second)
			blockingStep <- 1
		}
	}()
	return ch
}

func another_generator(arg string) <-chan Message {
	ch := make(chan Message)
	blockingStep := make(chan int)

	container, err := uDummy.ResolvedContainer()
	if err != nil {
		panic(err)
	}

	go container.Invoke(func(s sDummy.Service) {
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("Another: %s %d", arg, i), blockingStep}
			s.GetItem()
			blockingStep <- 1
		}
	})

	return ch
}
