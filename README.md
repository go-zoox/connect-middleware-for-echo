# Connect Middleware for echo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/connect-middleware-for-echo)](https://pkg.go.dev/github.com/go-zoox/connect-middleware-for-echo)
[![Build Status](https://github.com/go-zoox/connect-middleware-for-echo/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/connect-middleware-for-echo/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/connect-middleware-for-echo)](https://goreportcard.com/report/github.com/go-zoox/connect-middleware-for-echo)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/connect-middleware-for-echo/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/connect-middleware-for-echo?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/connect-middleware-for-echo.svg)](https://github.com/go-zoox/connect-middleware-for-echo/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/connect-middleware-for-echo.svg?label=Release)](https://github.com/go-zoox/connect-middleware-for-echo/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/connect-middleware-for-echo
```

## Getting Started

```go
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
```

## Related Projects
* [go-zoox/connect](https://github.com/go-zoox/connect) - The Auth Connect.
* [go-zoox/connect-middleware-for-zoox](https://github.com/go-zoox/connect-middleware-for-zoox) - The Auth Connect Middleware for Zoox.

## License
GoZoox is released under the [MIT License](./LICENSE).
