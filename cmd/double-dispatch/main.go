package main

import (
	"fmt"

	"github.com/maguroguma/go-experimental/internal/model/student"
	"github.com/maguroguma/go-experimental/internal/model/subject"
)

func main() {
	undergraduateStudent := student.NewUndergraduateStudent(20)
	masterStudent := student.NewMasterStudent("Shushi Taro")
	doctorStudent := student.NewDoctorStudent(27, "Hakase Jiro")

	liberalArts := subject.NewLiberalArts()
	quantumMechanics := subject.NewQuantumMechanics()
	graduationResearch := subject.NewGraduationResearch()

	fmt.Println(
		"under graduate student can register liberal arts: ", canRegister(undergraduateStudent, liberalArts),
	)
	fmt.Println(
		"master student can register liberal arts: ", canRegister(masterStudent, liberalArts),
	)
	fmt.Println(
		"doctor student can register liberal arts: ", canRegister(doctorStudent, liberalArts),
	)
	fmt.Println(
		"under graduate student can register quantum mechanics: ", canRegister(undergraduateStudent, quantumMechanics),
	)
	fmt.Println(
		"master student can register quantum mechanics: ", canRegister(masterStudent, quantumMechanics),
	)
	fmt.Println(
		"doctor student can register quantum mechanics: ", canRegister(doctorStudent, quantumMechanics),
	)
	fmt.Println(
		"under graduate student can register graduation research: ", canRegister(undergraduateStudent, graduationResearch),
	)
	fmt.Println(
		"master student can register graduation research: ", canRegister(masterStudent, graduationResearch),
	)
	fmt.Println(
		"doctor student can register graduation research: ", canRegister(doctorStudent, graduationResearch),
	)

	fmt.Println("===")

	fmt.Println("under graduate student liberal arts grade:", calculateGrade(undergraduateStudent, liberalArts))
	fmt.Println("master student liberal arts grade:", calculateGrade(masterStudent, liberalArts))
	fmt.Println("doctor student liberal arts grade:", calculateGrade(doctorStudent, liberalArts))
	fmt.Println("under graduate student quantum mechanics grade:", calculateGrade(undergraduateStudent, quantumMechanics))
	fmt.Println("master student quantum mechanics grade:", calculateGrade(masterStudent, quantumMechanics))
	fmt.Println("doctor student quantum mechanics grade:", calculateGrade(doctorStudent, quantumMechanics))
	fmt.Println("under graduate student graduation research grade:", calculateGrade(undergraduateStudent, graduationResearch))
	fmt.Println("master student graduation research grade:", calculateGrade(masterStudent, graduationResearch))
	fmt.Println("doctor student graduation research grade:", calculateGrade(doctorStudent, graduationResearch))
}

func canRegister(st student.Student, su student.Subject) bool {
	return st.CanRegister(su)
}

func calculateGrade(st student.Student, su student.Subject) int {
	return st.CalculateGrade(su)
}
