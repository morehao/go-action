package main

import (
	"fmt"

	"github.com/morehao/go-action/bizcase/gotag/enterprise"
	"github.com/morehao/go-action/bizcase/gotag/impl"
	"github.com/morehao/go-action/bizcase/gotag/platform"
)

func main() {
	oss := impl.NewOSS()
	auth := impl.NewAuth()
	audit := impl.NewAuditLog()
	report := enterprise.NewReportService()

	p := platform.NewPlatform(oss, auth, audit, report)
	p.Run()

	fmt.Println("\n=== Testing Features ===")

	token, err := auth.Login("admin", "admin")
	if err != nil {
		fmt.Printf("Auth failed: %v\n", err)
	} else {
		fmt.Printf("Login success, token: %s\n", token)
	}

	err = oss.Upload("my-bucket", "hello.txt", []byte("Hello World"))
	if err != nil {
		fmt.Printf("Upload failed: %v\n", err)
	} else {
		fmt.Println("Upload success")
	}

	audit.Log("file.upload", "admin")

	reportContent, err := report.GenerateReport("Monthly Summary")
	if err != nil {
		fmt.Printf("Report failed: %v\n", err)
	} else {
		fmt.Println(reportContent)
	}
}
