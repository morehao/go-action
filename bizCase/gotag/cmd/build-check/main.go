// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	projectDir, _ := filepath.Abs(filepath.Join("..", ".."))
	fmt.Println("Checking project:", projectDir)

	checks := []struct {
		name string
		cmd  *exec.Cmd
	}{
		{
			name: "OpenSource build (no tags)",
			cmd:  exec.Command("go", "build", "-o", "/dev/null", "."),
		},
		{
			name: "Enterprise build (-tags enterprise)",
			cmd:  exec.Command("go", "build", "-tags", "enterprise", "-o", "/dev/null", "."),
		},
	}

	hasError := false
	for _, check := range checks {
		check.cmd.Dir = projectDir
		fmt.Printf("\n=== %s ===\n", check.name)
		output, err := check.cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("FAIL: %v\n%s\n", err, output)
			hasError = true
		} else {
			fmt.Println("PASS")
		}
	}

	if hasError {
		fmt.Println("\n❌ Some checks failed")
		os.Exit(1)
	}
	fmt.Println("\n✅ All checks passed")
}
