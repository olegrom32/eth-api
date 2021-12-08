package rest

import (
	"github.com/gofiber/fiber/v2"

	"github.com/olegrom32/tech-assignment/application/handler/groupinfo"
)

// GroupInfoResponse is the HTTP response
type GroupInfoResponse struct {
	Name    string   `json:"name"`
	Indexes []uint64 `json:"indexes"`
}

// GroupInfo handles group info request
// @Summary Get group info
// @Description Finds a group by ID and returns it
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} GroupInfoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups/{id} [get]
func (c *Controller) GroupInfo(ctx *fiber.Ctx) error {
	groupID, err := ctx.ParamsInt("id")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(ErrorResponse{Message: err.Error()})
	}

	g, err := c.groupInfoHandler.Handle(groupinfo.NewCommand(int64(groupID)))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)

		return ctx.JSON(ErrorResponse{Message: err.Error()})
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(GroupInfoResponse{
		Name:    g.Name,
		Indexes: g.Indexes,
	})
}
