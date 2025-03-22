package student

import "github.com/maguroguma/go-experimental/internal/model/subject"

type Student interface {
	CanRegister(subject.Subject) bool
}

type UndergraduateStudent struct {
}

func NewUndergraduateStudent() *UndergraduateStudent {
	return &UndergraduateStudent{}
}

func (u *UndergraduateStudent) CanRegister(s subject.Subject) bool {
	return s.CanRegisterUndergraduateStudent()
}

type MasterStudent struct {
}

func NewMasterStudent() *MasterStudent {
	return &MasterStudent{}
}

func (m *MasterStudent) CanRegister(s subject.Subject) bool {
	return s.CanRegisterMasterStudent()
}
