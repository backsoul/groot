package controllers

import (
	"strings"

	"github.com/backsoul/groot/configs"
	"github.com/backsoul/groot/pkg/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Me(ctx *fiber.Ctx) error {
	accessToken := strings.ReplaceAll(string(ctx.Get("Authorization")), "Bearer ", "")
	token, err := jwt.ParseWithClaims(accessToken, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Get("JWT_KEY")), nil
	})
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Errot parse jwt",
			"data":    err.Error(),
		})
	}
	if token.Valid {
		// user := types.User{}
		// result := database.DB().First(&user, "email = ?", payload.Email)
		return ctx.JSON(fiber.Map{
			"status": "data",
			"data":   token.Claims,
		})
	} else {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Errot token invalid",
			"data":    err.Error(),
		})
	}
}
