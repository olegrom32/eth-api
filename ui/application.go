package ui

import (
	"github.com/gofiber/fiber/v2"

	"github.com/olegrom32/tech-assignment/config"
	"github.com/olegrom32/tech-assignment/ui/rest"
)

type Application struct {
	restServer *fiber.App
	restAddr   string
}

func NewApplication(
	restServer *fiber.App,
	restCfg config.RESTServer,
	_ *rest.Controller,
) *Application {
	return &Application{
		restServer: restServer,
		restAddr:   restCfg.Addr,
	}
}

func (a *Application) Run() error {
	// Use run.Group to run multiple servers like HTTP, gRPC, Machinery works etc.
	return a.restServer.Listen(a.restAddr)
}
