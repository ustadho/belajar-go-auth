package models

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	CompanyID string   `json:"companyId,omitempty`
	Username  string   `json:"username,omitempty`
	Roles     []string `json:"roles,omitempty`
	jwt.StandardClaims
}

const ip = "192.168.0.107"

func (claims *JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(ip, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}
