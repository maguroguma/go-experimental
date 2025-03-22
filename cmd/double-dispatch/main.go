package main

import (
	"fmt"

	"github.com/maguroguma/go-experimental/internal/model/student"
	"github.com/maguroguma/go-experimental/internal/model/subject"
)

func main() {
	undergraduateStudent := student.NewUndergraduateStudent()
	masterStudent := student.NewMasterStudent()

	liberalArts := subject.NewLiberalArts()
	quantumMechanics := subject.NewQuantumMechanics()
	graduationResearch := subject.NewGraduationResearch()

	fmt.Println(
		"under graduate student can register liberal arts: ", canRegister(undergraduateStudent, liberalArts),
	)
	fmt.Println(
		"under graduate student can register quantum mechanics: ", canRegister(undergraduateStudent, quantumMechanics),
	)
	fmt.Println(
		"master student can register liberal arts: ", canRegister(masterStudent, liberalArts),
	)
	fmt.Println(
		"master student can register quantum mechanics: ", canRegister(masterStudent, quantumMechanics),
	)
	fmt.Println(
		"under graduate student can register graduation research: ", canRegister(undergraduateStudent, graduationResearch),
	)
	fmt.Println(
		"master student can register graduation research: ", canRegister(masterStudent, graduationResearch),
	)
}

func canRegister(st student.Student, su subject.Subject) bool {
	return st.CanRegister(su)
}
