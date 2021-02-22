package helpers

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ConvertMetaData(c *fiber.Ctx) (uint, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(float64)
	str := ParseFloat(id)
	uID, err := ParseUint(str)
	if err != nil {
		return uID, err
	}
	return uID, nil
}
