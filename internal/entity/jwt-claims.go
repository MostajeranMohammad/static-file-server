package entity

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserId   string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}
