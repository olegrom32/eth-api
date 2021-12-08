package group

import (
	"github.com/olegrom32/tech-assignment/domain/shared"
)

type Model struct {
	ID      shared.GroupID
	Name    Name
	Indexes []shared.IndexID
}
