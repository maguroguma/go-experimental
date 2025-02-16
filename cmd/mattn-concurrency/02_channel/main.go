package main

import (
	"fmt"
	"time"
)

func main() {
	// チャンネルは make で作る
	// chan はデフォルトだとバッファサイズが1
	ch := make(chan string)
	go channel_print("Hello", ch)
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
