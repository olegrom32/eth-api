package rest

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// BlockLatest handles latest block request
// @Summary Get the latest block
// @Description Finds the latest block and returns it
// @Accept json
// @Produce json
// @Success 200 {object} BlockInfoResponse
// @Failure 500 {object} ErrorResponse
// @Router /blocks/latest [get]
func (c *Controller) BlockLatest(ctx *fiber.Ctx) error {
	b, err := c.blockInfoHandler.HandleLatest(ctx.Context())
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)

		return ctx.JSON(ErrorResponse{Message: err.Error()})
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(BlockInfoResponse{
		Number:            b.Number,
		Hash:              b.Hash,
		GasUsed:           b.GasUsed,
		Timestamp:         b.ReceivedAt.Format(time.RFC3339),
		TransactionsCount: b.TransactionsCount,
	})
}
