package memo

import "github.com/maguroguma/go-experimental/internal/model/memo"

//go:generate mockgen -destination=mock_memo.go -package=memo github.com/maguroguma/go-experimental/internal/service/memo Repository,Service

type Repository interface {
	Find(id int) (*memo.Memo, error)
}

type Service interface {
	GetMemo(id int) (*memo.Memo, error)
}
