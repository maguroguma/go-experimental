package main

import (
	"fmt"
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	uDummy "github.com/maguroguma/go-experimental/internal/usecase/dummy"
)

func main() {
	// ch := generator("Hello")
	ch := another_generator()

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

// 受信専用のチャネルは generator と見立てることが出来る
func generator(msg string) <-chan string {
	ch := make(chan string)

	// 無名ゴルーチン
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()

	return ch // 双方向のチャネルは受信専用としても振る舞うことができる
}

func another_generator() <-chan string {
	ch := make(chan string)

	container, err := uDummy.ResolvedContainer()
	if err != nil {
		panic(err)
	}

	go container.Invoke(func(s sDummy.Service) {
		for i := 0; ; i++ {
			msg, _ := s.GetItem()
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	})

	return ch
}
