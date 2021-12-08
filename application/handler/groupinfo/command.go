package groupinfo

import (
	"github.com/olegrom32/tech-assignment/domain/shared"
)

type Command struct {
	groupID shared.GroupID
}

func NewCommand(groupID int64) Command {
	return Command{
		groupID: shared.NewGroupID(groupID),
	}
}
