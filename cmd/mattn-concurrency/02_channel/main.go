package main

import (
	"fmt"
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	uDummy "github.com/maguroguma/go-experimental/internal/usecase/dummy"
)

func main() {
	// チャンネルは make で作る
	// chan はデフォルトだとバッファサイズが1
	ch := make(chan string)

	// go channel_print("Hello", ch)
	go another_channel_print(ch)

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch) // 消費するまでブロックされる
	}
	fmt.Println("Done!")
}

func channel_print(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}

func another_channel_print(ch chan string) {
	c, err := uDummy.ResolvedContainer()
	if err != nil {
		panic(err)
	}

	c.Invoke(func(s sDummy.Service) {
		for i := 0; ; i++ {
			msg, _ := s.GetItem()
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	})
}
