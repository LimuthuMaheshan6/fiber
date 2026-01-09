package auth

import (
	"log"
	

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserRoutes(router fiber.Router) {

	router.Get("/login",  func (c *fiber.Ctx) error{
		claims := jwt.MapClaims{}
		claims["name"] = "limuthu"
		claims["age"] = 22
	
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenGenerated, err := token.SignedString([]byte("secret"))
		
		if err != nil {
			log.Fatal(err, "jwt generate error")
		}

		log.Println("Claims\t",tokenGenerated)
		return c.SendString("Login Works")
	})

	router.Post("/login", func (c *fiber.Ctx) error {

		claims := jwt.MapClaims{}
		claims["name"] = "limuthu"
		claims["age"] = 22
	
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenGenerated, err := token.SignedString([]byte("secret"))
		
		if err != nil {
			log.Fatal(err, "jwt generate error")
		}

		log.Println("Claims\t",tokenGenerated)

		c.Cookie(&fiber.Cookie{
			Name: "Bearer",
			Value: tokenGenerated,
		})

		return c.JSON(fiber.Map{
			"status" : "Login Successful",
		})

		
	})


}

