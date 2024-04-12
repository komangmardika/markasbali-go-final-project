package middlewares

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

func CheckAuth(ctx *fiber.Ctx) error {
	au := string(ctx.Request().Header.Peek("Authorization"))
	aus := strings.Split(au, " ")

	decodedBytes, err := base64.StdEncoding.DecodeString(aus[1])

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(map[string]any{
			"message": "access unauthorized",
		})
	}

	// Convert bytes to string
	decodedString := string(decodedBytes)
	if decodedString == os.Getenv("SECRET_TOKEN") {
		return ctx.Next()
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(map[string]any{
		"message": "access unauthorized",
	})
}
