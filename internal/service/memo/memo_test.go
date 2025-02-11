package memo_test

import (
	"errors"
	"testing"

	dMemo "github.com/maguroguma/go-experimental/internal/model/memo"
	sMemo "github.com/maguroguma/go-experimental/internal/service/memo"
	"go.uber.org/dig"
	"go.uber.org/mock/gomock"
)

func TestService_GetMemo(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := sMemo.NewMockRepository(ctrl)

	mockRepo.EXPECT().Find(gomock.Eq(99)).DoAndReturn(func(_ int) (*dMemo.Memo, error) {
		return &dMemo.Memo{ID: 99, Text: "TEST"}, nil
	}).AnyTimes()
	mockRepo.EXPECT().Find(gomock.Eq(1)).Return(nil, errors.New("Not found")).Times(1)

	s := sMemo.NewService(mockRepo)

	m, err := s.GetMemo(99)
	t.Log(m, err)
	m, err = s.GetMemo(99)
	t.Log(m, err)
	m, err = s.GetMemo(99)
	t.Log(m, err)
	if err != nil {
		t.Fatal(err)
	}
	if !(m.ID == 99 && m.Text == "TEST") {
		t.Fatal("unexpected")
	}

	m2, err2 := s.GetMemo(1)
	t.Log(m2, err2)
	if err2 == nil {
		t.Fatal("unexpected")
	}
}

func TestService_GetMemo_UsingDig(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := sMemo.NewMockRepository(ctrl)
	mockRepo.EXPECT().Find(gomock.Eq(99)).DoAndReturn(func(_ int) (*dMemo.Memo, error) {
		return &dMemo.Memo{ID: 99, Text: "TEST"}, nil
	}).AnyTimes()
	mockRepo.EXPECT().Find(gomock.Eq(1)).Return(nil, errors.New("Not found")).Times(1)

	// 設定済みの mock repository を返す
	mockRepoProvider := func() sMemo.Repository {
		return mockRepo
	}

	container := dig.New()
	container.Provide(sMemo.NewService)
	container.Provide(mockRepoProvider)
	container.Invoke(func(s sMemo.Service) {
		m, err := s.GetMemo(99)
		t.Log(m, err)
		m, err = s.GetMemo(99)
		t.Log(m, err)
		m, err = s.GetMemo(99)
		t.Log(m, err)
		if err != nil {
			t.Fatal(err)
		}
		if !(m.ID == 99 && m.Text == "TEST") {
			t.Fatal("unexpected")
		}

		m2, err2 := s.GetMemo(1)
		t.Log(m2, err2)
		if err2 == nil {
			t.Fatal("unexpected")
		}
	})
}
