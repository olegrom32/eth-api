package index

import (
	"errors"
)

var ErrInvalidUSDPrice = errors.New("invalid usd price")

type USDPrice uint64

func NewUSDPrice(value int64) (USDPrice, error) {
	if value <= 0 {
		return 0, ErrInvalidUSDPrice
	}

	return USDPrice(value), nil
}

func (vo USDPrice) Uint64() uint64 {
	return uint64(vo)
}
