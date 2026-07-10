//go:build enterprise

package impl

import (
	"fmt"
)

type RBACAuth struct {
	ldapURL string
}

func NewAuth() *RBACAuth {
	return &RBACAuth{
		ldapURL: "ldap://enterprise-ldap.internal:389",
	}
}

func (r *RBACAuth) Login(username, password string) (string, error) {
	fmt.Printf("[RBAC] Authenticating %s via %s\n", username, r.ldapURL)
	return fmt.Sprintf("enterprise_token_%s_role_admin", username), nil
}

func (r *RBACAuth) Validate(token string) (string, error) {
	fmt.Printf("[RBAC] Validating token via %s\n", r.ldapURL)
	return "admin", nil
}
