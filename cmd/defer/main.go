package main

import "fmt"

func main() {
	test1(true)
	fmt.Println("===")
	test1(false)
	fmt.Println("===")
	testWhenPanic()
	fmt.Println("===")
	methodChain()
	fmt.Println("===")
	howParameter()
	fmt.Println("===")
	b()
	fmt.Println("===")

	fmt.Println("overwrite result:", overwrite())
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

func overwrite() (result int) {
	defer func() {
		result = 42 // return の値を書き換える
	}()
	return 1
}

func testWhenPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered in testWhenPanic: %v\n", r)
		}
	}()

	panic("panic in testWhenPanic")
}

func methodChain() {
	fmt.Println("BEGIN: methodChain")

	m := &MyStruct{}
	defer m.Start().DoSomething().End()

	fmt.Println("END: methodChain")
}

type MyStruct struct{}

func (m *MyStruct) Start() *MyStruct {
	fmt.Println("Start")
	return m
}

func (m *MyStruct) DoSomething() *MyStruct {
	fmt.Println("DoSomething")
	return m
}

func (m *MyStruct) End() *MyStruct {
	fmt.Println("End")
	return m
}

func howParameter() {
	double := func(x int) int {
		fmt.Println("double called:", x)
		return x * 2
	}
	plus := func(x, y int) int {
		fmt.Println("plus called:", x, y)
		return x + y
	}

	fmt.Println("howParameter START")
	defer plus(plus(10, 100), double(1))
	fmt.Println("howParameter END")
}

// ref: https://go.dev/doc/effective_go#defer
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
