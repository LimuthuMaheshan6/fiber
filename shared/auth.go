package shared

import "github.com/gofiber/fiber/v2"

func Auth(c *fiber.Ctx) error{

	var token string

	if token == ""{

		return c.Status(404).JSON(fiber.Map {
		   "error": "Doesn't Work",
		})

	}

	return c.Next();
}
