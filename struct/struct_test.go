package _struct

import (
	"fmt"
	"testing"
)

func Test_NestedStruct(t *testing.T) {
	NestedStruct()
}

func Test_InterviewExercises1(t *testing.T) {
	InterviewExercises1()
}

func Test_InterviewExercises2(t *testing.T) {
	InterviewExercises2()
}

func Test_NewPerson(t *testing.T) {
	person := NewPerson("pprof.cn", "北京", 90)
	fmt.Println(person)
}

func Test_Inherit(t *testing.T) {
	Inherit()
}
