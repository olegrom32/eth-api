package indexinfo

import (
	"github.com/olegrom32/tech-assignment/domain/index"
	"github.com/olegrom32/tech-assignment/domain/shared"
)

type IndexRepository interface {
	FindByID(shared.IndexID) (index.Model, error)
}

type Response struct {
	Name              string
	ETHPriceInWei     uint64
	USDPriceInCents   uint64
	USDCapitalization uint64
	PercentageChange  int64
}

type Handler struct {
	indexRepository IndexRepository
}

func NewHandler(indexRepository IndexRepository) *Handler {
	return &Handler{
		indexRepository: indexRepository,
	}
}

func (h *Handler) Handle(cmd Command) (Response, error) {
	data, err := h.indexRepository.FindByID(cmd.indexID)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Name:              data.Name.String(),
		ETHPriceInWei:     data.ETHPriceInWei.Uint64(),
		USDPriceInCents:   data.USDPriceInCents.Uint64(),
		USDCapitalization: data.USDCapitalization.Uint64(),
		PercentageChange:  data.PercentageChange.Int64(),
	}, nil
}
