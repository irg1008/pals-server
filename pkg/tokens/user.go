package tokens

import (
	"fmt"
	"irg1008/next-go/ent"
	"irg1008/next-go/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
)

const (
	accessTokenScope  = "access_token"
	refreshTokenScope = "refresh_token"
	userKey           = "user"
)

type Payload struct {
	Email string `json:"email" mapstructure:"email"`
}

type RefreshClaims struct {
	Id int `json:"id"`
}

type TokenPair struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func UnsignUser(t *jwt.Token) (user *Payload, err error) {
	claims := t.Claims.(jwt.MapClaims)
	payload := claims[userKey]
	err = mapstructure.Decode(payload, &user)
	return
}

func (s *Signing) CreateUserToken(user *ent.User) (token string, err error) {
	var payload Payload
	err = copier.Copy(&payload, &user)
	if err != nil {
		return
	}

	claims := jwt.MapClaims{userKey: payload}
	baseClaims := BaseClaims{user.ID, config.TokenDuration, accessTokenScope}

	return s.SignWithClaims(claims, baseClaims)
}

func (s *Signing) CreateUserRefreshToken(user *ent.User) (string, error) {
	baseClaims := BaseClaims{user.ID, config.RefreshTokenDuration, refreshTokenScope}
	return s.Sign(baseClaims)
}

func (s *Signing) CerateUserTokenPair(user *ent.User) (tokens TokenPair, err error) {
	token, err := s.CreateUserToken(user)
	if err != nil {
		return
	}

	refresh, err := s.CreateUserRefreshToken(user)
	if err != nil {
		return
	}

	return TokenPair{token, refresh}, nil
}

func (s *Signing) ParseRefreshToken(token string) (*RefreshClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if validAlg := token.Method.Alg() == SigningAlgorithm().Alg(); !validAlg {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	claims := t.Claims.(jwt.MapClaims)
	scope := claims["scope"].(string)

	if scope != refreshTokenScope {
		return nil, fmt.Errorf("Invalid token scope, used scope: %v, expected scope: %v", scope, refreshTokenScope)
	}

	resfreshClaims := &RefreshClaims{
		Id: int(claims["sub"].(float64)),
	}
	return resfreshClaims, nil
}
