package shared

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Name string `json:"name"`
	Age int `json:"age"`
	jwt.RegisteredClaims 
}


func Auth(c *fiber.Ctx) error{

	

	
    
	token :=  c.Cookies("jwt_access")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map {
			"error": "Missing Token",
		})
	}

	log.Println("cookie:", token)


/* 	var jwtKey = []byte("secret")
	
	claims := &jwt.RegisteredClaims{}
	
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
 */	

     _, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte("secret"), nil
	 })

	if err != nil {
		fmt.Println("Token is invalid: ", err)
		return c.JSON(fiber.Map {"Status": "Wrong jwt",})
	}

	

	return c.Next()
	
	
}



