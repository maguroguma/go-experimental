package dummy

//go:generate mockgen -destination=mock_dummy.go -package=dummy github.com/maguroguma/go-experimental/internal/service/dummy Client,Service

type Client interface {
	Get() (string, error)
}

type Service interface {
	GetItem() (string, error)
}
