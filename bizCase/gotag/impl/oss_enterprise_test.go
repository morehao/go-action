//go:build enterprise

package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliyunOSS_Upload(t *testing.T) {
	oss := NewOSS()
	err := oss.Upload("bucket", "key", []byte("data"))
	assert.NoError(t, err)
}

func TestRBACAuth_Login(t *testing.T) {
	auth := NewAuth()
	token, err := auth.Login("admin", "pass")
	assert.NoError(t, err)
	assert.Contains(t, token, "enterprise_token")
}