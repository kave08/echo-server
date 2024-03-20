package security

import "github.com/golang-jwt/jwt"

type JtwClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
