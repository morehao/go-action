//go:build !enterprise

package impl

type StubAuditLog struct{}

func NewAuditLog() *StubAuditLog {
	return &StubAuditLog{}
}

func (s *StubAuditLog) Log(action, user string) {
}
