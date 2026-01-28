package contact

import "github.com/gofiber/fiber/v2"

func ContactRouter(router fiber.Router) {
	Contact(router)
}
