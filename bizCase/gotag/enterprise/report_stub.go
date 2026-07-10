//go:build !enterprise

package enterprise

import (
	"fmt"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (r *ReportService) GenerateReport(title string) (string, error) {
	return "", fmt.Errorf("report is only available in enterprise edition")
}
