package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "missing token"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := jwt.MapClaims{}
		jwtSecret := os.Getenv("JWT_SECRET")

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid or expired token"})
		}

		c.Set("user", claims)

		fmt.Println("JWT claims:", claims)

		return next(c)
	}
}

func RequireRole(requireRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user")
			claims := user.(*jwt.MapClaims)

			role := (*claims)["role"]
			if role != requireRole {
				return c.JSON(http.StatusForbidden, echo.Map{"message": "you are not allowed to access this resource"})
			}

			fmt.Println("Role is token", role)

			return next(c)
		}
	}
}
