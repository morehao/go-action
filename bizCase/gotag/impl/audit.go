//go:build enterprise

package impl

import (
	"fmt"
	"time"
)

type ESAuditLog struct {
	esEndpoint string
}

func NewAuditLog() *ESAuditLog {
	return &ESAuditLog{
		esEndpoint: "http://elasticsearch.internal:9200",
	}
}

func (e *ESAuditLog) Log(action, user string) {
	fmt.Printf("[Audit][ES] time=%s action=%s user=%s endpoint=%s\n",
		time.Now().Format(time.RFC3339), action, user, e.esEndpoint)
}
