package controllers

import (
	"strings"

	"github.com/backsoul/groot/configs"
	"github.com/backsoul/groot/pkg/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Refresh(ctx *fiber.Ctx) error {
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
	claims := token.Claims.(*types.UserClaims)
	NewToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JwtSecret := configs.Get("JWT_KEY")
	tokenJwt, err := NewToken.SignedString([]byte(JwtSecret))
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  "error",
			"message": "Error SignedString",
			"data":    err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"accessToken":  accessToken,
			"refreshToken": tokenJwt,
		},
	})
}
