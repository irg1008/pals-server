package tokens

import (
	"irg1008/next-go/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func baseClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"exp": time.Now().Add(config.TokenDuration).Unix(),
	}
}

func newToken() (*jwt.Token, jwt.MapClaims) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, baseClaims())
	return token, token.Claims.(jwt.MapClaims)
}

func newTokenWithClaims(claims jwt.MapClaims) *jwt.Token {
	for k, v := range baseClaims() {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

type Signing struct {
	secret string
}

func NewSigning(secret string) *Signing {
	return &Signing{
		secret: secret,
	}
}
func (s *Signing) Sign(token *jwt.Token) (string, error) {
	return token.SignedString([]byte(s.secret))
}

func (s *Signing) SignWithClaims(claims jwt.MapClaims) (string, error) {
	token := newTokenWithClaims(claims)
	return s.Sign(token)
}
