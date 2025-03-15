package main

import "fmt"

func main() {
	test1(true)
	fmt.Println("===")
	test1(false)
}

func test1(isSuspend bool) {
	defer func() {
		fmt.Println("test1 defer 1")
	}()
	defer func() {
		fmt.Println("test1 defer 2")
	}()

	fmt.Println("test1 body")

	if isSuspend {
		return
	}

	defer func() {
		fmt.Println("test1 defer 3")
	}()
}
