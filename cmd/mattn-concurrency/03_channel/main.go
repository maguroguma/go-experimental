package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator("Hello")
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
