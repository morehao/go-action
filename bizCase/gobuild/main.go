package main

import (
	"fmt"

	"github.com/morehao/go-action/bizCase/gobuild/version"
)

func main() {
	fmt.Println("Version:", version.GetDeployVersion())
	fmt.Println("Mode:", version.GetDeployMode())
}
