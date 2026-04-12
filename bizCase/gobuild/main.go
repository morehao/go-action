package main

import (
	"fmt"

	"github.com/morehao/go-action/bizcase/gobuild/version"
)

func main() {
	fmt.Println("Version:", version.GetDeployVersion())
	fmt.Println("Mode:", version.GetDeployMode())
}
