package rest

import (
	"github.com/gofiber/fiber/v2"

	"github.com/olegrom32/tech-assignment/application/handler/indexinfo"
)

type IndexInfoResponse struct {
	Name              string `json:"name"`
	ETHPriceInWei     uint64 `json:"eth_price_in_wei"`
	USDPriceInCents   uint64 `json:"usd_price_in_cents"`
	USDCapitalization uint64 `json:"usd_capitalization"`
	PercentageChange  int64  `json:"percentage_change"`
}

// IndexInfo handles the index info request
// @Summary Get index info
// @Description Finds an index by ID and returns it
// @Accept json
// @Produce json
// @Param id path string true "Index ID"
// @Success 200 {object} IndexInfoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /indexes/{id} [get]
func (c *Controller) IndexInfo(ctx *fiber.Ctx) error {
	indexID, err := ctx.ParamsInt("id")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(ErrorResponse{Message: err.Error()})
	}

	data, err := c.indexInfoHandler.Handle(indexinfo.NewCommand(int64(indexID)))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)

		return ctx.JSON(ErrorResponse{Message: err.Error()})
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(IndexInfoResponse{
		Name:              data.Name,
		ETHPriceInWei:     data.ETHPriceInWei,
		USDPriceInCents:   data.USDPriceInCents,
		USDCapitalization: data.USDCapitalization,
		PercentageChange:  data.PercentageChange,
	})
}
