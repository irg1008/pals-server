package services

import (
	"errors"
	"irg1008/next-go/ent"
	"irg1008/next-go/ent/authrequest"
	"irg1008/next-go/ent/user"
	"irg1008/next-go/pkg/crypt"
	"irg1008/next-go/pkg/db"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	DB *db.DB
}

func (s *AuthService) GetUserByEmail(email string) (*ent.User, error) {
	return s.DB.User.Query().Where(user.EmailEQ(email)).Only(s.DB.Ctx)
}

func (s *AuthService) CreateUser(email string, password string) (*ent.User, error) {
	hash, err := crypt.Hash(password)

	if err != nil {
		return nil, err
	}

	return s.DB.User.Create().SetEmail(email).SetPassword(hash).Save(s.DB.Ctx)
}

func (s *AuthService) GetUserById(id int) (*ent.User, error) {
	return s.DB.User.Get(s.DB.Ctx, id)
}

func (s *AuthService) deleteRequests(userID int, requestType authrequest.Type) error {
	isUserRequest := authrequest.HasUserWith(user.ID(userID))
	isSameType := authrequest.TypeEQ(requestType)
	_, err := s.DB.AuthRequest.Delete().
		Where(isUserRequest, isSameType).
		Exec(s.DB.Ctx)
	return err
}

const (
	confirmExpiryTime = 15 * time.Minute
)

type CreateRequestOptions struct {
	invalidateRest bool
}

func (s *AuthService) createRequest(
	email string,
	requestType authrequest.Type,
	opts CreateRequestOptions,
) (*ent.AuthRequest, error) {
	requestUser, err := s.DB.User.Query().Where(user.EmailEQ(email)).Only(s.DB.Ctx)

	if requestUser == nil {
		return nil, err
	}

	isUserRequest := authrequest.HasUserWith(user.ID(requestUser.ID))
	isSameType := authrequest.TypeEQ(requestType)
	requestUpdate := s.DB.AuthRequest.Update().Where(isUserRequest, isSameType)

	if opts.invalidateRest {
		requestUpdate = requestUpdate.SetActive(false)
	}

	if err := requestUpdate.Exec(s.DB.Ctx); err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(confirmExpiryTime)
	req, err := s.DB.AuthRequest.Create().
		SetUser(requestUser).
		SetType(requestType).
		SetActive(true).
		SetExpiresAt(expiresAt).
		Save(s.DB.Ctx)

	if err != nil {
		return nil, err
	}

	return req, nil
}

type ConfirmRequestOptions struct {
	deleteAll bool
}

func (s *AuthService) confirmRequest(
	token string,
	requestType authrequest.Type,
	opts ConfirmRequestOptions,
) (*ent.User, error) {
	id, err := uuid.Parse(token)
	if err != nil {
		return nil, err
	}

	isSameToken := authrequest.TokenEQ(id)
	isSameType := authrequest.TypeEQ(requestType)
	req, err := s.DB.AuthRequest.Query().Where(isSameToken, isSameType).WithUser().Only(s.DB.Ctx)

	if err != nil {
		return nil, errors.New("request not found")
	}

	if !req.Active {
		return nil, errors.New("request is not active")
	}

	if req.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("request is expired")
	}

	if opts.deleteAll {
		err = s.deleteRequests(req.Edges.User.ID, requestType)
	} else {
		err = req.Update().SetActive(false).Exec(s.DB.Ctx)
	}

	if err != nil {
		return nil, err
	}

	return req.Edges.User, nil
}

func (s *AuthService) CreateConfirmationRequest(email string) (*ent.AuthRequest, error) {
	opts := CreateRequestOptions{invalidateRest: true}
	return s.createRequest(email, authrequest.TypeConfirmEmail, opts)
}

func (s *AuthService) ConfirmEmail(token string) (*ent.User, error) {
	opts := ConfirmRequestOptions{deleteAll: true}
	user, err := s.confirmRequest(token, authrequest.TypeConfirmEmail, opts)

	if err != nil {
		return nil, err
	}

	err = user.Update().SetIsConfirmed(true).Exec(s.DB.Ctx)
	return user, err
}

func (s *AuthService) CreatePasswordResetRequest(email string) (*ent.AuthRequest, error) {
	opts := CreateRequestOptions{invalidateRest: false}
	return s.createRequest(email, authrequest.TypeResetPassword, opts)
}

func (s *AuthService) ResetPassword(token string, password string) (*ent.User, error) {
	opts := ConfirmRequestOptions{deleteAll: true}
	user, err := s.confirmRequest(token, authrequest.TypeResetPassword, opts)

	if err != nil {
		return nil, err
	}

	hash, err := crypt.Hash(password)
	if err != nil {
		return nil, err
	}

	if err := user.Update().SetPassword(hash).Exec(s.DB.Ctx); err != nil {
		return nil, err
	}

	return user, nil
}
