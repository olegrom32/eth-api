package provider

import (
	"github.com/gofiber/fiber/v2"
)

func NewRESTServer() *fiber.App {
	return fiber.New()
}
