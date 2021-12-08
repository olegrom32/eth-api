package indexinfo

import (
	"github.com/olegrom32/tech-assignment/domain/shared"
)

type Command struct {
	indexID shared.IndexID
}

func NewCommand(indexID int64) Command {
	return Command{
		indexID: shared.NewIndexID(indexID),
	}
}
