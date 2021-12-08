//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"github.com/google/wire"

	"github.com/olegrom32/tech-assignment/application/handler/blockinfo"
	"github.com/olegrom32/tech-assignment/application/handler/groupinfo"
	"github.com/olegrom32/tech-assignment/application/handler/indexinfo"
	"github.com/olegrom32/tech-assignment/application/handler/listgroups"
	"github.com/olegrom32/tech-assignment/config"
	"github.com/olegrom32/tech-assignment/config/provider"
	"github.com/olegrom32/tech-assignment/infrastructure/repository"
	"github.com/olegrom32/tech-assignment/ui"
	"github.com/olegrom32/tech-assignment/ui/rest"
)

func build() (*ui.Application, error) {
	wire.Build(
		configSet,
		uiSet,
		infrastructureSet,
		applicationSet,

		ui.NewApplication,
	)

	return new(ui.Application), nil
}

var configSet = wire.NewSet(
	provider.NewRESTServer,
	provider.NewETHClient,
	provider.NewContract,
	config.NewConfig,
	wire.FieldsOf(
		new(*config.Config),
		"RESTServer",
		"Ropsten",
	),
)

var uiSet = wire.NewSet(
	rest.NewController,
)

var infrastructureSet = wire.NewSet(
	repository.NewBlock,
	wire.Bind(new(blockinfo.BlockRepository), new(*repository.Block)),

	repository.NewGroup,
	wire.Bind(new(listgroups.GroupRepository), new(*repository.Group)),
	wire.Bind(new(groupinfo.GroupRepository), new(*repository.Group)),

	repository.NewIndex,
	wire.Bind(new(indexinfo.IndexRepository), new(*repository.Index)),
)

var applicationSet = wire.NewSet(
	listgroups.NewHandler,
	wire.Bind(new(rest.ListGroupsHandler), new(*listgroups.Handler)),

	groupinfo.NewHandler,
	wire.Bind(new(rest.GroupInfoHandler), new(*groupinfo.Handler)),

	indexinfo.NewHandler,
	wire.Bind(new(rest.IndexInfoHandler), new(*indexinfo.Handler)),

	blockinfo.NewHandler,
	wire.Bind(new(rest.BlockInfoHandler), new(*blockinfo.Handler)),
)

func Initialize() (*ui.Application, error) {
	app, err := build()
	if err != nil {
		// log errors
		return nil, err
	}

	return app, nil
}
