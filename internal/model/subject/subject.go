package subject

type Subject interface {
	CanRegisterUndergraduateStudent() bool
	CanRegisterMasterStudent() bool
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
