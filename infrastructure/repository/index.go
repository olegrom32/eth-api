package repository

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"

	"github.com/olegrom32/tech-assignment/domain/index"
	"github.com/olegrom32/tech-assignment/domain/shared"
	"github.com/olegrom32/tech-assignment/pkg/contract"
)

type Index struct {
	client *contract.ContractCaller
}

func NewIndex(client *contract.ContractCaller) *Index {
	return &Index{
		client: client,
	}
}

func (r *Index) FindByID(indexID shared.IndexID) (index.Model, error) {
	data, err := r.client.GetIndex(&bind.CallOpts{}, big.NewInt(indexID.Int64()))
	if err != nil {
		return index.Model{}, errors.WithStack(err)
	}

	m := index.Model{
		PercentageChange: index.NewPercentageChange(data.PercentageChange.Int64()),
	}

	if m.Name, err = index.NewName(data.Name); err != nil {
		return index.Model{}, errors.WithStack(err)
	}

	if m.ETHPriceInWei, err = index.NewETHPrice(data.EthPriceInWei.Int64()); err != nil {
		return index.Model{}, errors.WithStack(err)
	}

	if m.USDPriceInCents, err = index.NewUSDPrice(data.UsdPriceInCents.Int64()); err != nil {
		return index.Model{}, errors.WithStack(err)
	}

	if m.USDCapitalization, err = index.NewUSDCapitalization(data.UsdCapitalization.Int64()); err != nil {
		return index.Model{}, errors.WithStack(err)
	}

	return m, nil
}
