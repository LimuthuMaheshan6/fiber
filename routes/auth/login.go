package auth

import (
	"log"
	"project/shared"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"


)


func UserRoutes(router fiber.Router) {


	router.Get("/login", shared.Auth,  func (c *fiber.Ctx) error{
		return c.JSON(fiber.Map {"status": "Login Works"})
	})

	router.Post("/login", func (c *fiber.Ctx) error {

		

		

		
		tokenGenerated, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"name": "Limuthu Maheshan", "age": 19}).SignedString([]byte("secret"))
		
		if err != nil {
			log.Fatal(err, "jwt generate error")
		}

		
		c.Cookie(&fiber.Cookie{
			Name: "jwt_access",
			Value: tokenGenerated,
			Expires: time.Now().Add(time.Minute*15),
			HTTPOnly: true,
			Secure: true,
			SameSite: "Lax",
			
		})
		c.Set("Set-Cookie", "")
		return c.JSON(fiber.Map{
					"status": "Login Successful...",
		})
	

		/* c.Set("x-auth-header", tokenGenerated)

		log.Print("jwt works")

		return c.JSON(fiber.Map {"Status": "Works Successfully"}) */

		
	})


}

