package controllers

import (
	"encoding/json"

	"github.com/backsoul/groot/internal/database"
	"github.com/backsoul/groot/pkg/types"
	"github.com/gofiber/fiber/v2"
)

func Me(ctx *fiber.Ctx) error {
	var payload map[string]interface{}
	json.Unmarshal([]byte(ctx.Body()), &payload)

	payloadJson, _ := json.Marshal(payload["user"])
	var user types.UserClaims
	json.Unmarshal([]byte(payloadJson), &user)
	User := types.User{}
	database.DB().Where("Email = ?", user.Email).First(&User)
	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   User,
	})
}
