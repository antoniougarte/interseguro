package config

import "os"

type JWTConfig struct {
	Secret     string
	Expiration int
}

func LoadJWT() *JWTConfig {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "mi-super-secreto-cambiar-en-produccion"
	}

	return &JWTConfig{
		Secret:     secret,
		Expiration: 24,
	}
}
