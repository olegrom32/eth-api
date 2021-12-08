package index

import (
	"errors"
)

var ErrInvalidUSDCapitalization = errors.New("invalid usd capitalization")

type USDCapitalization uint64

func NewUSDCapitalization(value int64) (USDCapitalization, error) {
	if value <= 0 {
		return 0, ErrInvalidUSDCapitalization
	}

	return USDCapitalization(value), nil
}

func (vo USDCapitalization) Uint64() uint64 {
	return uint64(vo)
}
