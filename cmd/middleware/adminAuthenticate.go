package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AdminAuthenticate(c *fiber.Ctx) error {
	type adminLogin struct {
		Username string `json:"identity"`
		Password string `json:"password"`
	}
	var (
		admin adminLogin
	)
	fmt.Print(admin)
	return nil
}
