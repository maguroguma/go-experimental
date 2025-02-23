package dummy

import (
	"time"

	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
)

func NewClient() sDummy.Client {
	return &client{}
}

type client struct{}

func (c *client) Get() (string, error) {
	// 時間がかかる処理
	time.Sleep(1 * time.Second)

	return "Hello, World from DUMMY API!", nil
}
