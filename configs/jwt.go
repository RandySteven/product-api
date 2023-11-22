package configs

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("cCM7rrsrmle6")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
