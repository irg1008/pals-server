package tokens

import (
	"irg1008/next-go/ent"

	"github.com/golang-jwt/jwt/v5"
)

type SignedUser struct {
	Email string
}

func UnsignUser(t *jwt.Token) *SignedUser {
	claims := t.Claims.(jwt.MapClaims)
	user := &SignedUser{
		Email: claims["email"].(string),
	}
	return user
}

func (s *Signing) SignUserToken(user *ent.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
	}
	return s.SignWithClaims(claims)
}
