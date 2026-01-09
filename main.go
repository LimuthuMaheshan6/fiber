package main

import (
	"log"
	"project/routes/auth"
	

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()


	auth.AuthApi(app)
	
	



	log.Println("Server Works...")
	err := app.Listen(":8000")
	if err != nil {
		log.Print("Server Failed... main.go ln-15 \t", err)
	}

	

}

