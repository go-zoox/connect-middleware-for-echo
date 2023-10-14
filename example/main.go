package main

import (
	"github.com/go-zoox/connect-middleware-for-echo"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()

	r.Use(connect.Create("YOUR_SECRET_KEY"))

	r.GET("/user", func(c echo.Context) error {
		user, err := connect.GetUser(c)
		if err != nil {
			return c.JSON(401, map[string]any{
				"message": "unauthorized",
			})
		}

		return c.JSON(200, map[string]any{
			"user": user,
		})
	})

	r.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]any{
			"message": "helloworld",
		})
	})

	r.Logger.Fatal(r.Start(":8080"))
}
