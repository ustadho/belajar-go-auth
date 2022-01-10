package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ustadho/belajar-go-auth/models"
)

const (
	jwtPrivateToken = "secret"
	ip              = "192.168.0.107"
)

func GenerateToken(claims *models.JwtClaims, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encode token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwtPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
