package main

import (
	"fmt"

	rMemo "github.com/maguroguma/go-experimental/internal/repository/memo"
	sMemo "github.com/maguroguma/go-experimental/internal/service/memo"
)

func main() {
	fmt.Println("Run...")

	r := rMemo.NewRepository()
	s := sMemo.NewService(r)
	m, err := s.GetMemo(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("id: %d, text: %s\n", m.ID, m.Text)

	fmt.Println("Done.")
}
