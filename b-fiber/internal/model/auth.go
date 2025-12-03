package model

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
