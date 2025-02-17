package main

import (
	"fmt"
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	uDummy "github.com/maguroguma/go-experimental/internal/usecase/dummy"
)

func main() {
	// ch := fanIn(generator("Hello"), generator("Bye"))
	// ch := fanIn(another_generator("Hello"), another_generator("Bye"))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }

	// select の実験
	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-another_generator("Hello"):
			fmt.Println(msg1)
			// case msg2 := <-another_generator("Bye"):
		case msg2 := <-generator("Hello"): // 毎回別の chan を作っている！
			fmt.Println(msg2)
		}
	}
}

// select は複数の chan の受信を同時に待てる
func fanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
			case s := <-ch2:
				new_ch <- s
			}
		}
	}()
	return new_ch
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("%s %d", msg, i)
			// time.Sleep(1 * time.Second) // 待機が後半だと受信側で消費した瞬間切断される形になる
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
