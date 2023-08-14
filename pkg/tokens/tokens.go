package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type BaseClaims struct {
	id       int
	duration time.Duration
	scope    string
}

type Signing struct {
	secret string
}

func mixClaims(claims jwt.MapClaims, baseClaims jwt.MapClaims) jwt.MapClaims {
	for k, v := range baseClaims {
		claims[k] = v
	}
	return claims
}

func mixWithBasicClaims(claims jwt.MapClaims, base BaseClaims) jwt.MapClaims {
	baseClaims := jwt.MapClaims{
		"exp":   time.Now().Add(base.duration).Unix(),
		"iat":   time.Now().Unix(),
		"sub":   base.id,
		"scope": base.scope,
		// TODO: We need to create a unique token id so we only have 1 valid refresh token at any moment. We could create a session table to store and be able t validate and invalidate refresh tokens at will
	}
	return mixClaims(claims, baseClaims)
}

func (s *Signing) SignWithClaims(claims jwt.MapClaims, base BaseClaims) (string, error) {
	claims = mixWithBasicClaims(claims, base)
	alg := SigningAlgorithm()
	token := jwt.NewWithClaims(alg, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *Signing) Sign(base BaseClaims) (string, error) {
	return s.SignWithClaims(jwt.MapClaims{}, base)
}

func NewSigning(secret string) *Signing {
	return &Signing{secret}
}

func SigningAlgorithm() jwt.SigningMethod {
	return jwt.SigningMethodHS256
}
