package blockinfo

import (
	"context"
	"time"

	"github.com/olegrom32/tech-assignment/domain/block"
)

type BlockRepository interface {
	FindByNumber(context.Context, block.Number) (block.Model, error)
	FindByHash(context.Context, block.Hash) (block.Model, error)
	FindLatest(context.Context) (block.Model, error)
}

type Response struct {
	Number            uint64
	Hash              string
	GasUsed           uint64
	ReceivedAt        time.Time
	TransactionsCount int
}

type Handler struct {
	blockRepository BlockRepository
}

func NewHandler(blockRepository BlockRepository) *Handler {
	return &Handler{
		blockRepository: blockRepository,
	}
}

func (h *Handler) HandleByNumber(ctx context.Context, cmd NumberCommand) (Response, error) {
	b, err := h.blockRepository.FindByNumber(ctx, cmd.number)
	if err != nil {
		return Response{}, err
	}

	return response(b), nil
}

func (h *Handler) HandleByHash(ctx context.Context, cmd HashCommand) (Response, error) {
	b, err := h.blockRepository.FindByHash(ctx, cmd.hash)
	if err != nil {
		return Response{}, err
	}

	return response(b), nil
}

func (h *Handler) HandleLatest(ctx context.Context) (Response, error) {
	b, err := h.blockRepository.FindLatest(ctx)
	if err != nil {
		return Response{}, err
	}

	return response(b), nil
}

func response(b block.Model) Response {
	return Response{
		Number:            b.Number.Uint64(),
		Hash:              b.Hash.String(),
		GasUsed:           b.GasUsed.Uint64(),
		ReceivedAt:        b.ReceivedAt,
		TransactionsCount: b.TransactionsCount.Int(),
	}
}
