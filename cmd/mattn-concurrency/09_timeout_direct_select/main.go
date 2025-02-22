package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator("Hi!")
	timeout := time.After(5 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout: // このチャンネルはずっと生きている！
			fmt.Println("5s Timeout!")
			return
		}
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
