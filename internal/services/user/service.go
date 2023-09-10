package user

import (
	"irg1008/pals/ent"
	"irg1008/pals/ent/userdata"
	"irg1008/pals/pkg/db"

	"github.com/go-pkgz/auth/token"
)

type UserService struct {
	DB *db.DB
}

func (s *UserService) GetUserData(id string) (*ent.UserData, error) {
	return s.DB.UserData.Query().Where(userdata.AuthID(id)).First(s.DB.Ctx)
}

func (s *UserService) CreateUserData(user *token.User) (*ent.UserData, error) {
	return s.DB.UserData.
		Create().
		SetAuthID(user.ID).
		SetName(user.Name).
		Save(s.DB.Ctx)
}

func (s *UserService) GetOrCreteUserData(user *token.User) (*ent.UserData, error) {
	data, err := s.GetUserData(user.ID)
	if err == nil {
		return data, nil
	}

	return s.CreateUserData(user)
}
