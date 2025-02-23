package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	ch := generator("Hi!", quit)
	for i := rand.Intn(5); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- true // main が先に終わる generator 側の出力がなされなこともある
}

func generator(msg string, quit chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg): // select では受信だけでなく送信も混ぜることができる
				// nothing
			case <-quit:
				fmt.Println("Goroutine done")
				return
			}
		}
	}()
	return ch
}
