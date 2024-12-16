package utils

import (
	"fmt"
	"testing"
)

func TestGetFirstLetter(t *testing.T) {
	t.Log(GetFirstLetter("你好"))
	tpl := "你好%s, %s"
	t.Log(fmt.Sprintf(tpl, "a", "b"))
	t.Log(fmt.Sprintf(tpl, "a"))
}
