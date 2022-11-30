package middleware

import (
	"fmt"
	"strings"

	"github.com/KaioMarxDEV/gofinance/cmd/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authorize(c *fiber.Ctx) error {
	jwtToken := c.Get("Authorization")

	_, tokenString, _ := strings.Cut(jwtToken, "Bearer ")

	if tokenString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(handler.ResponseHTTP{
			Success: false,
			Message: "Missing auth token",
			Data:    nil,
		})
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("SECRET"), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(handler.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("invalid token: %v", err),
			Data:    nil,
		})
	}

	fmt.Println(token)
	return c.Next()
}
