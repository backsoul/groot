package controllers

import (
	"github.com/backsoul/groot/internal/database"
	"github.com/backsoul/groot/pkg/types"
	"github.com/backsoul/groot/pkg/utils"
	"github.com/backsoul/groot/pkg/validators"
	"github.com/gofiber/fiber/v2"
)

func ControllerLogin(c *fiber.Ctx) error {
	var payload types.PayloadLoginUser
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
	user := types.User{}
	result := database.DB().First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   "user not found",
		})
	}
	passwordCorrect := utils.CheckPasswordHash(payload.Password, user.Password)
	if passwordCorrect == true {
		return c.JSON(fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"ID":        user.ID,
				"FirstName": user.FirstName,
				"LastName":  user.LastName,
				"Email":     user.Email,
			},
		})
	} else {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   "Password incorrect",
		})
	}
}
