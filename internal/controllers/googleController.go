package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/shareed2k/goth_fiber"
)

func ControllerAuthCallback(ctx *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.Send([]byte(user.Email))
}
