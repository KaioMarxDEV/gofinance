package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/montanaflynn/jsonwebtoken"
)

func Authenticate(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")

	_, string, _ := strings.Cut(bearerToken, "Bearer ")
	tokenString := strings.TrimSpace(string)

	token, err := jsonwebtoken.DecodeHMAC(tokenString, []byte("SECRET"))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(token)
	}

	return c.Next()
}
