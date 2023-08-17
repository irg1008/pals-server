package roles

import (
	"irg1008/next-go/pkg/tokens"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func userExtractionPanic() {
	slog.Error("Trying to extract user from context, but no auth middleware was assigned to route.")
	panic("Provide a valid middleware before calling this function.")
}

func userMappingPanic() {
	slog.Error("Failed to map user from token.")
	panic("Review token creation and mapping. Check all fields are assigned correctly.")
}

func GetUser(c echo.Context) *tokens.Payload {
	token := c.Get(contextKey)
	if token == nil {
		userExtractionPanic()
	}

	user, err := tokens.UnsignUser(token.(*jwt.Token))
	if err != nil {
		userMappingPanic()
	}

	return user

}