package service

import (
	"errors"
	"time"

	"github.com/antoniougarte/b-fiber/config"
	"github.com/antoniougarte/b-fiber/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

var jwtConfig = config.LoadJWT()

// ValidateCredentials - En producción, validar contra base de datos
func ValidateCredentials(username, password string) bool {
	// Por ahora, credenciales hardcodeadas para el demo
	validUsers := map[string]string{
		"admin": "admin123",
		"user":  "user123",
	}

	if validPassword, exists := validUsers[username]; exists {
		return validPassword == password
	}
	return false
}

// GenerateJWT genera un token JWT
func GenerateJWT(username string) (string, error) {
	claims := model.JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConfig.Expiration) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

// ValidateJWT valida un token JWT
func ValidateJWT(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
