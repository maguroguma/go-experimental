package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web    = fakeSearch("web")
	Web2   = fakeSearch("web")
	Image  = fakeSearch("image")
	Image2 = fakeSearch("image")
	Video  = fakeSearch("video")
	Video2 = fakeSearch("video")
)

type Result string
type Search func(query string) Result

func main() {
	// rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang") // collate results
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

// 時間のかかる関数(Search 型)を返す
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// query パラメータを使って複数の Search 関数を並行に走らせる
// そのうち、最初に結果をチャネルに送ったもののみ採用して返す
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web, Web2) }()
	go func() { c <- First(query, Image, Image2) }()
	go func() { c <- First(query, Video, Video2) }()

	timeout := time.After(80 * time.Millisecond)

	// fan-in に近いことをしている？
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}
