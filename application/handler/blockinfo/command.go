package blockinfo

import (
	"github.com/olegrom32/tech-assignment/domain/block"
)

type NumberCommand struct {
	number block.Number
}

type HashCommand struct {
	hash block.Hash
}

func NewNumberCommand(number uint64) NumberCommand {
	return NumberCommand{
		number: block.NewNumber(number),
	}
}

func NewHashCommand(hash string) HashCommand {
	return HashCommand{
		hash: block.NewHash(hash),
	}
}
