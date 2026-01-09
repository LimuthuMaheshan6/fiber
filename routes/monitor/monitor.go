package monitor

import (
	"github.com/gofiber/fiber/v2"
)

func MonitorWorkingReq(router fiber.Router) {


	router.Get("/monitor",  func (c *fiber.Ctx)  error{

		return c.SendString("Monitor Works")
		
	})

}