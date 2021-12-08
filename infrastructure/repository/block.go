package repository

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/olegrom32/tech-assignment/domain/block"
)

type Block struct {
	client *ethclient.Client
}

func NewBlock(client *ethclient.Client) *Block {
	return &Block{
		client: client,
	}
}

func (r *Block) FindByNumber(ctx context.Context, number block.Number) (block.Model, error) {
	b, err := r.client.BlockByNumber(ctx, big.NewInt(number.Int64()))
	if err != nil {
		return block.Model{}, errors.WithStack(err)
	}

	return convert(b), nil
}

func (r *Block) FindByHash(ctx context.Context, hash block.Hash) (block.Model, error) {
	b, err := r.client.BlockByHash(ctx, common.HexToHash(hash.String()))
	if err != nil {
		return block.Model{}, errors.WithStack(err)
	}

	return convert(b), nil
}

func (r *Block) FindLatest(ctx context.Context) (block.Model, error) {
	n, err := r.client.BlockNumber(ctx)
	if err != nil {
		return block.Model{}, errors.WithStack(err)
	}

	return r.FindByNumber(ctx, block.NewNumber(n))
}

func convert(b *types.Block) block.Model {
	return block.Model{
		Number:            block.NewNumber(b.Number().Uint64()),
		Hash:              block.NewHash(b.Hash().String()),
		GasUsed:           block.NewGasUsed(b.GasUsed()),
		ReceivedAt:        b.ReceivedAt,
		TransactionsCount: block.NewTransactionsCount(b.Transactions().Len()),
	}
}
