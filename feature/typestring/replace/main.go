package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  你 好 "
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.TrimSpace(strings.Replace(str, " ", "", -1)))
}
