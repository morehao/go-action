package big

import (
	"fmt"
	"testing"
)

func TestSetPrec(t *testing.T) {
	res := SetPrec(3.33333333, 100)
	fmt.Println("res:", res)
}
