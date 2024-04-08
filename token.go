package connect

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// GetToken get user from context.
func GetToken(ctx echo.Context) (token string, err error) {
	token = ctx.Request().Header.Get("x-connect-token")
	if token == "" {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}

// MustGetToken get user from context.
func MustGetToken(ctx echo.Context) (token string) {
	token, err := GetToken(ctx)
	if err != nil {
		panic(err)
	}

	return token
}
