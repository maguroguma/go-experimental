package subject

type Subject interface {
	CanRegisterUndergraduateStudent() bool
	CanRegisterMasterStudent() bool
	CanRegisterDoctorStudent() bool
}

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
