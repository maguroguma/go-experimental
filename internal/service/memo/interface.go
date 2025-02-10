package memo

import "github.com/maguroguma/go-experimental/internal/model/memo"

type Repository interface {
	Find(id int) (*memo.Memo, error)
}

type Service interface {
	GetMemo(id int) (*memo.Memo, error)
}
