package security

import "github.com/golang-jwt/jwt"

type JtwClaims struct {
	UserName string `json:"username"`
	UserId   string `json:"user_id"`
	jwt.StandardClaims
}
