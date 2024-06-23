package middleware

import (
	"assignment/config"
	"net/http"

	"assignment/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func SetHeaders(c echo.Context) {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("X-XSS-Protection", "1; mode=block")
	c.Response().Header().Set("X-Content-Type-Options", "nosniff")
	c.Response().Header().Set("X-Frame-Options", "DENY")
}

func SecurityMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		next(c)
		SetHeaders(c)
		return nil
	}
}

// AuthMiddleware checks if the request is authenticated
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		config := config.Init("config")
		token := c.Request().Header.Get("Authorization")
		claims := jwt.MapClaims{}
		_, parseErr := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("jwt_secret.key")), nil
		})

		if parseErr != nil {
			return c.JSON(http.StatusForbidden, utils.GenericResponse{Message: "Forbidden"})
		}
		return next(c)
	}
}
