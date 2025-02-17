package unitTest

import "testing"

func TestAdd(t *testing.T) {
	var (
		a        = 1
		b        = 2
		expected = 4
	)
	res := Add(a, b)
	if res != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, res, expected)
	}
}
