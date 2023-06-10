package typeArray

import (
	"fmt"
	"testing"
)

func Test_test1(t *testing.T) {
	test1()
}

func Test_findTargetIndex(t *testing.T) {
	fmt.Printf("result:%v", findTargetIndex([]int{1, 3, 5, 8, 7}, 8))
}

func Test_declare(t *testing.T) {
	declare()
}
