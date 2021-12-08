package groupinfo

import (
	"github.com/olegrom32/tech-assignment/domain/group"
	"github.com/olegrom32/tech-assignment/domain/shared"
)

type GroupRepository interface {
	FindByID(shared.GroupID) (group.Model, error)
}

type Response struct {
	Name    string
	Indexes []uint64
}

type Handler struct {
	groupRepository GroupRepository
}

func NewHandler(groupRepository GroupRepository) *Handler {
	return &Handler{
		groupRepository: groupRepository,
	}
}

func (h *Handler) Handle(cmd Command) (Response, error) {
	g, err := h.groupRepository.FindByID(cmd.groupID)
	if err != nil {
		return Response{}, err
	}

	idx := make([]uint64, len(g.Indexes))

	for k, v := range g.Indexes {
		idx[k] = v.Uint64()
	}

	return Response{
		Name:    g.Name.String(),
		Indexes: idx,
	}, nil
}
