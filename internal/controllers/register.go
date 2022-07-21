package controllers

import (
	"github.com/backsoul/groot/configs"
	"github.com/backsoul/groot/pkg/models"
	"github.com/backsoul/groot/pkg/types"
	"github.com/backsoul/groot/pkg/utils"
	"github.com/backsoul/groot/pkg/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ControllerRegister(c *fiber.Ctx) error {
	var payload types.PayloadRegisterUser
	if err := c.BodyParser(&payload); err != nil {
		c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}
	errors := validators.ValidateStruct(payload)
	if errors != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   errors,
		})
	}
	claims := types.UserClaims{
		FirstName:      payload.FirstName,
		LastName:       payload.LastName,
		Email:          payload.Email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JwtSecret := configs.Get("JWT_KEY")
	tokenJwt, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		panic(err)
	}
	password, _ := utils.HashPassword(payload.Password)
	_, err = models.CreateUser(payload.FirstName, payload.LastName, payload.Email, password)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "email already exists",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   tokenJwt,
	})
}
