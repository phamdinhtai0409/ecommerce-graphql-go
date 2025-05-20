package util

import (
	"os"
)

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "secret-key"
	}
	return secret
}

func GetJWTExpiration() string {
	expiration := os.Getenv("JWT_EXPIRATION")
	if expiration == "" {
		return "24h" // Default 24h
	}
	return expiration
}
