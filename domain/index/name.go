package index

import (
	"errors"
)

var ErrInvalidName = errors.New("invalid index name")

type Name string

func NewName(value string) (Name, error) {
	if value == "" {
		return "", ErrInvalidName
	}

	return Name(value), nil
}

func (vo Name) String() string {
	return string(vo)
}
