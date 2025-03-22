package subject

import "github.com/maguroguma/go-experimental/internal/model/student"

type LiberalArts struct {
}

func NewLiberalArts() *LiberalArts {
	return &LiberalArts{}
}
func (l *LiberalArts) CanRegisterUndergraduateStudent() bool {
	return true
}
func (l *LiberalArts) CanRegisterMasterStudent() bool {
	return false
}
func (l *LiberalArts) CanRegisterDoctorStudent() bool {
	return false
}
func (l *LiberalArts) CalculateUndergraduateStudentGrade(u *student.UndergraduateStudent) int {
	return u.Age * 1
}
func (l *LiberalArts) CalculateMasterStudentGrade(m *student.MasterStudent) int {
	return len(m.Name) * 1
}
func (l *LiberalArts) CalculateDoctorStudentGrade(d *student.DoctorStudent) int {
	return (d.Age + len(d.Name)) * 1
}

type QuantumMechanics struct {
}

func NewQuantumMechanics() *QuantumMechanics {
	return &QuantumMechanics{}
}
func (q *QuantumMechanics) CanRegisterUndergraduateStudent() bool {
	return false
}
func (q *QuantumMechanics) CanRegisterMasterStudent() bool {
	return true
}
func (q *QuantumMechanics) CanRegisterDoctorStudent() bool {
	return false
}
func (q *QuantumMechanics) CalculateUndergraduateStudentGrade(u *student.UndergraduateStudent) int {
	return u.Age * 2
}
func (q *QuantumMechanics) CalculateMasterStudentGrade(m *student.MasterStudent) int {
	return len(m.Name) * 2
}
func (q *QuantumMechanics) CalculateDoctorStudentGrade(d *student.DoctorStudent) int {
	return (d.Age + len(d.Name)) * 2
}

type GraduationResearch struct {
}

func NewGraduationResearch() *GraduationResearch {
	return &GraduationResearch{}
}
func (g *GraduationResearch) CanRegisterUndergraduateStudent() bool {
	return true
}
func (g *GraduationResearch) CanRegisterMasterStudent() bool {
	return true
}
func (g *GraduationResearch) CanRegisterDoctorStudent() bool {
	return true
}
func (g *GraduationResearch) CalculateUndergraduateStudentGrade(u *student.UndergraduateStudent) int {
	return u.Age * 3

}
func (g *GraduationResearch) CalculateMasterStudentGrade(m *student.MasterStudent) int {
	return len(m.Name) * 3
}
func (g *GraduationResearch) CalculateDoctorStudentGrade(d *student.DoctorStudent) int {
	return (d.Age + len(d.Name)) * 3
}
