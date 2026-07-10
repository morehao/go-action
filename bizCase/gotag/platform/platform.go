package platform

import "fmt"

type OSS interface {
	Upload(bucket, key string, data []byte) error
	Download(bucket, key string) ([]byte, error)
}

type Auth interface {
	Login(username, password string) (string, error)
	Validate(token string) (string, error)
}

type AuditLog interface {
	Log(action, user string)
}

type Report interface {
	GenerateReport(title string) (string, error)
}

type Platform struct {
	OSS     OSS
	Auth    Auth
	Audit   AuditLog
	Report  Report
}

func NewPlatform(oss OSS, auth Auth, audit AuditLog, report Report) *Platform {
	return &Platform{OSS: oss, Auth: auth, Audit: audit, Report: report}
}

func (p *Platform) Run() {
	fmt.Println("=== Platform Running ===")
	fmt.Println("OSS:", p.OSS)
	fmt.Println("Auth:", p.Auth)
	fmt.Println("Audit:", p.Audit)
	fmt.Println("Report:", p.Report)
}
