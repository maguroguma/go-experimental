package memo

import (
	mMemo "github.com/maguroguma/go-experimental/internal/model/memo"
	sMemo "github.com/maguroguma/go-experimental/internal/service/memo"
)

func NewRepository() sMemo.Repository {
	return &repository{}
}

type repository struct{}

func (r *repository) Find(id int) (*mMemo.Memo, error) {
	return &mMemo.Memo{ID: id, Text: "Hello, World!"}, nil
}
