package auth

import (
	

	"github.com/gofiber/fiber/v2"
)


func AuthApi(router fiber.Router) {
	UserRoutes(router)
}
