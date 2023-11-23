package configs

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("cCM7rrsrmle6")

type JWTClaim struct {
	User *models.User
	jwt.RegisteredClaims
}
