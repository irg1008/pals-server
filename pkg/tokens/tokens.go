package tokens

import (
	"irg1008/next-go/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var baseClaims = jwt.MapClaims{
	"exp": time.Now().Add(48 * time.Hour).Unix(),
}

func New() (*jwt.Token, jwt.MapClaims) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, baseClaims)
	return token, token.Claims.(jwt.MapClaims)
}

func NewWithClaims(claims jwt.MapClaims) *jwt.Token {
	for k, v := range baseClaims {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func Sign(token *jwt.Token) (string, error) {
	return token.SignedString([]byte(config.Env.JWTSecret))
}

func SignWithClaims(claims jwt.MapClaims) (string, error) {
	token := NewWithClaims(claims)
	return Sign(token)
}
