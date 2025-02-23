package main

import (
	"fmt"
	"time"
)

func main() {
	// waitUntilClose()
	// waitSending()
	// notifyByClose()
	// forRangeChan()
	// selectChan()
	// waitMultiClosedChan()
	doubleClose()
}

// close した chan からはバッファ分まで受信できる
// すべて受信したあとはゼロ値を受信して終了する
func waitUntilClose() {
	ch := make(chan int, 3)
	done := make(chan int)

	// 1秒ごとに、3回送信する
	// 終わったら done で報告する
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
		close(ch)
		done <- 0
	}()

	<-done
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		} else {
			fmt.Println(v)
		}
	}
}

func waitSending() {
	ch := make(chan int, 1) // make(chan int) だとバッファなしになってしまう

	go func() {
		for i := 1; ; i++ {
			ch <- i
			fmt.Printf("No %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	// 受信せずに一定時間後に close する
	time.Sleep(3 * time.Second)
	close(ch)
	fmt.Printf("after close: %d\n", <-ch)
	time.Sleep(3 * time.Second)
}

func notifyByClose() {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()

	v, ok := <-ch
	fmt.Printf("Done! v: %d, ok: %v\n", v, ok)
}

func forRangeChan() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func selectChan() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second)
		close(ch)
	}()

	select {
	case v, ok := <-ch:
		fmt.Println(v, ok) // 0 false
	case v, ok := <-time.After(time.Second):
		fmt.Println("timeout", v, ok)
	}
}

// 複数の goroutine で chan の受信を確認する
func waitMultiClosedChan() {
	done := make(chan struct{})

	go func() {
		select {
		case v, ok := <-done:
			fmt.Println("done child 1", v, ok)
		}
	}()
	go func() {
		select {
		case v, ok := <-done:
			fmt.Println("done child 2", v, ok)
		}
	}()

	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("done parent")
}

// panic になる
func doubleClose() {
	ch := make(chan int)

	close(ch)
	close(ch)
}
