package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jayantodpuji/grocerfy/internal/delivery"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtKey string) echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey: []byte(jwtKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return delivery.ResponseError(c, http.StatusUnauthorized, err.Error())
		},
		SuccessHandler: func(c echo.Context) {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*jwt.RegisteredClaims)
			c.Set("user_id", claims.Subject)
		},
	}
	return echojwt.WithConfig(config)
}
