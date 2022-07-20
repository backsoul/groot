package controllers

import (
	"time"

	"github.com/backsoul/groot/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ControllerRegister(c *fiber.Ctx) error {
	type MyCustomClaims struct {
		Email    string `json:"email"`
		Birthday int64  `json:"birthday"`
		jwt.StandardClaims
	}

	// Create the claims
	claims := MyCustomClaims{
		"hello@friendsofgo.tech",
		time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC).Unix(),
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "Friends of Go",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	mySecret := configs.Get("JWT_KEY")
	signedToken, err := token.SignedString([]byte(mySecret))
	if err != nil {
		panic(err)
	}
	return c.SendString(signedToken)
}
