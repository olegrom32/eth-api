package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/olegrom32/tech-assignment/application"
)

// ListGroupsResponse is the HTTP response
type ListGroupsResponse struct {
	Groups []uint64 `json:"groups"`
}

// ListGroups handles the list groups request
// @Summary Get groups
// @Description Returns list of groups
// @Accept json
// @Produce json
// @Success 200 {object} ListGroupsResponse
// @Failure 404
// @Failure 500 {object} ErrorResponse
// @Router /groups [get]
func (c *Controller) ListGroups(ctx *fiber.Ctx) error {
	data, err := c.listGroupsHandler.Handle()
	if err != nil {
		switch errors.Cause(err) {
		case application.ErrNoGroupsFound:
			return ctx.SendStatus(fiber.StatusNotFound)
		default:
			ctx.Status(fiber.StatusInternalServerError)

			return ctx.JSON(ErrorResponse{Message: err.Error()})
		}
	}

	resp := ListGroupsResponse{
		Groups: make([]uint64, len(data.List)),
	}

	for k, v := range data.List {
		resp.Groups[k] = v
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(resp)
}
