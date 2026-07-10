//go:build !enterprise

package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocalStorage_UploadAndDownload(t *testing.T) {
	oss := NewOSS()
	err := oss.Upload("test-bucket", "test-key", []byte("hello"))
	require.NoError(t, err)

	data, err := oss.Download("test-bucket", "test-key")
	require.NoError(t, err)
	assert.Equal(t, "hello", string(data))
}

func TestJWTAuth_Login(t *testing.T) {
	auth := NewAuth()
	token, err := auth.Login("admin", "admin")
	require.NoError(t, err)
	assert.Contains(t, token, "token_")
}

func TestJWTAuth_Login_Invalid(t *testing.T) {
	auth := NewAuth()
	_, err := auth.Login("admin", "wrong")
	assert.Error(t, err)
}