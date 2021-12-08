package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/olegrom32/tech-assignment/application/handler/blockinfo"
	"github.com/olegrom32/tech-assignment/application/handler/groupinfo"
	"github.com/olegrom32/tech-assignment/application/handler/indexinfo"
	"github.com/olegrom32/tech-assignment/application/handler/listgroups"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type ListGroupsHandler interface {
	Handle() (listgroups.Response, error)
}

type GroupInfoHandler interface {
	Handle(groupinfo.Command) (groupinfo.Response, error)
}

type IndexInfoHandler interface {
	Handle(indexinfo.Command) (indexinfo.Response, error)
}

type BlockInfoHandler interface {
	HandleByNumber(context.Context, blockinfo.NumberCommand) (blockinfo.Response, error)
	HandleByHash(context.Context, blockinfo.HashCommand) (blockinfo.Response, error)
	HandleLatest(context.Context) (blockinfo.Response, error)
}

type Controller struct {
	listGroupsHandler ListGroupsHandler
	groupInfoHandler  GroupInfoHandler
	indexInfoHandler  IndexInfoHandler
	blockInfoHandler  BlockInfoHandler
}

func NewController(
	router *fiber.App,
	listGroupsHandler ListGroupsHandler,
	groupInfoHandler GroupInfoHandler,
	indexInfoHandler IndexInfoHandler,
	blockInfoHandler BlockInfoHandler,
) *Controller {
	c := &Controller{
		listGroupsHandler: listGroupsHandler,
		groupInfoHandler:  groupInfoHandler,
		indexInfoHandler:  indexInfoHandler,
		blockInfoHandler:  blockInfoHandler,
	}

	router.Get("/groups", c.ListGroups)
	router.Get("/groups/:id", c.GroupInfo)
	router.Get("/indexes/:id", c.IndexInfo)
	router.Get("/blocks/latest", c.BlockLatest)
	router.Get("/blocks/:id", c.BlockInfo)

	return c
}
