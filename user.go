package connect

import (
	"fmt"

	user "github.com/go-zoox/connect/user"
	"github.com/labstack/echo/v4"
)

// GetUser gets the user from the context.
func GetUser(ctx echo.Context) (u *user.User, err error) {
	v := ctx.Get(ContextUserKey)
	if v == nil {
		return nil, fmt.Errorf("user not found")
	}

	u, ok := v.(*user.User)
	if !ok {
		return nil, fmt.Errorf("user invalid")
	}

	return u, nil
}

// MustGetUser get user from context.
func MustGetUser(ctx echo.Context) (u *user.User) {
	u, err := GetUser(ctx)
	if err != nil {
		panic(err)
	}

	return u
}
