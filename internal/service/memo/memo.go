package memo

import "github.com/maguroguma/go-experimental/internal/model/memo"

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

type service struct {
	repository Repository
}

func (s *service) GetMemo(id int) (*memo.Memo, error) {
	memo, err := s.repository.Find(id)
	if err != nil {
		return nil, err
	}
	return memo, nil
}
