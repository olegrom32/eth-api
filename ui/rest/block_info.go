package rest

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/olegrom32/tech-assignment/application/handler/blockinfo"
)

type BlockInfoResponse struct {
	Number            uint64 `json:"number"`
	Hash              string `json:"hash"`
	GasUsed           uint64 `json:"gas_used"`
	Timestamp         string `json:"timestamp"`
	TransactionsCount int    `json:"transactions_count"`
}

// BlockInfo handles block info request
// @Summary Get block info by hash or by number
// @Description Finds a block by given hash or number and returns it
// @Accept json
// @Produce json
// @Param id path string true "Block Hash or Number"
// @Success 200 {object} BlockInfoResponse
// @Failure 500 {object} ErrorResponse
// @Router /blocks/{id} [get]
func (c *Controller) BlockInfo(ctx *fiber.Ctx) error {
	var (
		b    blockinfo.Response
		hErr error
	)

	number, err := ctx.ParamsInt("id")
	if err == nil {
		b, hErr = c.blockInfoHandler.HandleByNumber(ctx.Context(), blockinfo.NewNumberCommand(uint64(number)))
	} else {
		b, hErr = c.blockInfoHandler.HandleByHash(ctx.Context(), blockinfo.NewHashCommand(ctx.Params("id")))
	}

	if hErr != nil {
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
