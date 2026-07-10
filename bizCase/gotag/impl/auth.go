//go:build !enterprise

package impl

import (
	"fmt"
)

type JWTAuth struct {
	secret string
}

func NewAuth() *JWTAuth {
	return &JWTAuth{secret: "opensource-secret"}
}

func (j *JWTAuth) Login(username, password string) (string, error) {
	if username == "admin" && password == "admin" {
		return fmt.Sprintf("token_%s", username), nil
	}
	return "", fmt.Errorf("invalid credentials")
}

func (j *JWTAuth) Validate(token string) (string, error) {
	if len(token) > 0 {
		return "admin", nil
	}
	return "", fmt.Errorf("invalid token")
}
