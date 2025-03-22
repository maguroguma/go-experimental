package student

// Acceptor に相当する
// 各メソッドは AcceptXxx に相当する
type Student interface {
	CanRegister(Subject) bool
	CalculateGrade(Subject) int
}

// Visitor に相当する
// 各メソッドは VisitXxx に相当する
type Subject interface {
	CanRegisterUndergraduateStudent() bool
	CanRegisterMasterStudent() bool
	CanRegisterDoctorStudent() bool
	CalculateUndergraduateStudentGrade(*UndergraduateStudent) int
	CalculateMasterStudentGrade(*MasterStudent) int
	CalculateDoctorStudentGrade(*DoctorStudent) int
}

type UndergraduateStudent struct {
	Age int
}

func NewUndergraduateStudent(age int) *UndergraduateStudent {
	return &UndergraduateStudent{Age: age}
}

func (u *UndergraduateStudent) CanRegister(s Subject) bool {
	return s.CanRegisterUndergraduateStudent()
}
func (u *UndergraduateStudent) CalculateGrade(s Subject) int {
	return s.CalculateUndergraduateStudentGrade(u)
}

type MasterStudent struct {
	Name string
}

func NewMasterStudent(name string) *MasterStudent {
	return &MasterStudent{Name: name}
}

func (m *MasterStudent) CanRegister(s Subject) bool {
	return s.CanRegisterMasterStudent()
}
func (m *MasterStudent) CalculateGrade(s Subject) int {
	return s.CalculateMasterStudentGrade(m)
}

type DoctorStudent struct {
	Age  int
	Name string
}

func NewDoctorStudent(age int, name string) *DoctorStudent {
	return &DoctorStudent{Age: age, Name: name}
}

func (d *DoctorStudent) CanRegister(s Subject) bool {
	return s.CanRegisterDoctorStudent()
}
func (d *DoctorStudent) CalculateGrade(s Subject) int {
	return s.CalculateDoctorStudentGrade(d)
}
