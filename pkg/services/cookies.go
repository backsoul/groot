package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddNewCookie(name string, value string, expires time.Time) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	cookie.HTTPOnly = true
	cookie.Domain = "groot.backsoul.xyz"
	cookie.Path = "/"
	return cookie
}
