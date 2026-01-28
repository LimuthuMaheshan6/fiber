package contact

import "github.com/gofiber/fiber/v2"


func Contact(router fiber.Router) {

	router.Post("/contact", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map {
			"purpose": "Works",
		})
	})



	router.Get("/contact", func(c *fiber.Ctx) error{
		return c.JSON(fiber.Map {

		})
	})

	router.Put("/contact", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map {
			
		})
	})

	router.Delete("/contact", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map {
			
		})
	})





}
