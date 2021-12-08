package test

import (
	"context"
	"log"

	"gopkg.in/h2non/baloo.v3"

	"github.com/olegrom32/tech-assignment/config"
	"github.com/olegrom32/tech-assignment/config/wire"
	"github.com/olegrom32/tech-assignment/pkg/testing"
)

var (
	unit      *testing.Unit
	container *Container
)

type Container struct {
	Ctx    context.Context
	Config *config.Config

	RESTClient *baloo.Client
}

func init() {
	var err error

	unit = testing.DefaultUnit()
	container = &Container{
		Ctx:        unit.Ctx,
		RESTClient: unit.WithRESTClient(),
	}

	if container.Config, err = config.NewConfig(); err != nil {
		log.Fatal(err)
	}
}

func SetUp() (*Container, func()) {
	unit.WithInitApp(func() {
		app, err := wire.Initialize()
		if err != nil {
			log.Fatalln("can't init application: ", err)
		}

		go app.Run()
	})

	return container, unit.TearDown
}
