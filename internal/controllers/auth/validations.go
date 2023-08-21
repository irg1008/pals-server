package auth

import (
	"irg1008/pals/ent"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (u *AuthController) userExistForEmail(email string) (*ent.User, error) {
	user, err := u.service.GetUserByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "No user found with this email")
	}
	return user, nil
}

func (u *AuthController) userExists(id int) (*ent.User, error) {
	user, err := u.service.GetUserById(id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return user, nil
}

func isNotConfirmed(user *ent.User) error {
	if !user.IsConfirmed {
		return echo.NewHTTPError(http.StatusForbidden, "Confirm your email first")
	}
	return nil

}

func isConfirmed(user *ent.User) error {
	if user.IsConfirmed {
		return echo.NewHTTPError(http.StatusConflict, "User already confirmed")
	}
	return nil
}
