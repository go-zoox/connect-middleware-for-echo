package connect

import (
	user "github.com/go-zoox/connect/user"
	"github.com/go-zoox/jwt"
	"github.com/labstack/echo/v4"
)

// CreateOptions is the options for the Create middleware.
type CreateOptions struct {
	RequireAuth bool
}

// Create creates a connect middleware for gin.
func Create(secretKey string, opts ...*CreateOptions) echo.MiddlewareFunc {
	var signer jwt.Jwt
	var optsX *CreateOptions
	if len(opts) > 0 && opts[0] != nil {
		optsX = opts[0]
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if signer == nil {
				signer = jwt.New(secretKey)
			}

			token := ctx.Request().Header.Get("x-connect-token")
			if token != "" {
				user := &user.User{}
				if err := user.Decode(signer, token); err != nil {
					// if ctx.AcceptJSON() {
					// 	ctx.JSON(http.StatusUnauthorized, gin.H{
					// 		"code":    401001,
					// 		"message": "Unauthorized",
					// 	})
					// } else {
					// 	ctx.Status(401)
					// }

					return echo.ErrUnauthorized
				}

				ctx.Set(ContextUserKey, user)
			}

			if optsX != nil && optsX.RequireAuth {
				if v := ctx.Get(ContextUserKey); v == nil {
					// if ctx.AcceptJSON() {
					// 	ctx.JSON(http.StatusUnauthorized, gin.H{
					// 		"code":    401001,
					// 		"message": "Unauthorized",
					// 	})
					// } else {
					// 	ctx.Status(401)
					// }

					return echo.ErrUnauthorized
				}
			}

			return next(ctx)
		}
	}
}
