package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	errorIs()
	fmt.Println("===")
	errorAs()
}

func errorIs() {
	err := fmt.Errorf("wrap: %w", os.ErrNotExist)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("ファイルが存在しません")
	}
	if errors.Is(err, os.ErrPermission) {
		fmt.Println("ファイルのアクセス権限がありません")
	} else {
		fmt.Println("")
	}

	certainErr1 := errors.New("特定のエラー1")
	certainErr2 := errors.New("特定のエラー2")
	wrapErr1 := fmt.Errorf("wrap: %w", certainErr1)
	wrapErr2 := fmt.Errorf("wrap: %w", certainErr2)
	if errors.Is(wrapErr1, certainErr1) {
		fmt.Println("1: OK")
	}
	if errors.Is(wrapErr1, certainErr2) {
		fmt.Println("2: NG")
	}
	if errors.Is(wrapErr2, certainErr1) {
		fmt.Println("3: NG")
	}
	if errors.Is(wrapErr2, certainErr2) {
		fmt.Println("4: OK")
	}
}

type MyError struct {
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("エラーコード: %d", e.Code)
}

func errorAs() {
	err := fmt.Errorf("wrap: %w", &MyError{Code: 404})

	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("特定のエラーが発生: %d\n", myErr.Code)
	}
}
