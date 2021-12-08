package index

import (
	"errors"
)

var ErrInvalidETHPrice = errors.New("invalid eth price")

type ETHPrice uint64

func NewETHPrice(value int64) (ETHPrice, error) {
	if value <= 0 {
		return 0, ErrInvalidETHPrice
	}

	return ETHPrice(value), nil
}

func (vo ETHPrice) Uint64() uint64 {
	return uint64(vo)
}
