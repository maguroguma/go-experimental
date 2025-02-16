package dummy

import (
	sDummy "github.com/maguroguma/go-experimental/internal/service/dummy"
	wDummy "github.com/maguroguma/go-experimental/internal/web/dummy"
	"go.uber.org/dig"
)

func ResolvedContainer() (*dig.Container, error) {
	container := dig.New()

	if err := container.Provide(sDummy.NewService); err != nil {
		return nil, err
	}
	if err := container.Provide(wDummy.NewClient); err != nil {
		return nil, err
	}

	return container, nil
}
