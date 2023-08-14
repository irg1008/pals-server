package services

import (
	"irg1008/next-go/ent"
	"irg1008/next-go/ent/user"
	"irg1008/next-go/pkg/crypt"
	"irg1008/next-go/pkg/db"
)

type AuthService struct {
	DB *db.DB
}

func (u *AuthService) GetUserByEmail(email string) (*ent.User, error) {
	return u.DB.User.Query().Where(user.EmailEQ(email)).Only(u.DB.Ctx)
}

func (u *AuthService) CreateUser(email string, password string) (*ent.User, error) {
	hash, err := crypt.Hash(password)

	if err != nil {
		return nil, err
	}

	return u.DB.User.Create().SetEmail(email).SetPassword(hash).Save(u.DB.Ctx)
}

func (u *AuthService) GetUserById(id int) (*ent.User, error) {
	return u.DB.User.Get(u.DB.Ctx, id)
}
