package middlewares

import (
	"os"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func jwtAuth() func(ctx *fiber.Ctx) {
	jwtSecret := os.Getenv("JWT_KEY")
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(jwtSecret),
	})
}
