package oop

import (
	"fmt"
	"testing"
)

func Test_Speak(t *testing.T) {
	var peo People
	a := Student{}
	peo = &a
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
