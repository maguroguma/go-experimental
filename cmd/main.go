package main

import (
	"fmt"

	rMemo "github.com/maguroguma/go-experimental/internal/repository/memo"
	sMemo "github.com/maguroguma/go-experimental/internal/service/memo"
	"go.uber.org/dig"
)

func main() {
	fmt.Println("Run...")
	fmt.Println("")

	fmt.Println("simple()")
	simple()
	fmt.Println("diByDig()")
	diByDig()

	fmt.Println("Done.")
}

func run(s sMemo.Service, l Logger) {
	m, err := s.GetMemo(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := fmt.Sprintf("id: %d, text: %s\n", m.ID, m.Text)
	l.Log(str)
}

func simple() {
	r := rMemo.NewRepository()
	s := sMemo.NewService(r)
	l := NewLogger()

	run(s, l)
}

func diByDig() {
	container := dig.New()

	// Provide する順序は何でも良い
	if err := container.Provide(NewLogger); err != nil {
		fmt.Println(err)
		return
	}
	if err := container.Provide(sMemo.NewService); err != nil {
		fmt.Println(err)
		return
	}
	if err := container.Provide(rMemo.NewRepository); err != nil {
		fmt.Println(err)
		return
	}

	if err := container.Invoke(run); err != nil {
		fmt.Println(err)
		return
	}
}
