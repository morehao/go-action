package main

import (
	"fmt"
	"testing"
)

func TestSetPrec(t *testing.T) {
	a, b := 2.3329, 3.1234
	// a, b := 2.33, 3.12
	n := a + b
	fmt.Println("n:", n)
	res := SetPrec(n, 2)
	fmt.Println("res:", res)
}
