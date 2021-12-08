package listgroups

import (
	"github.com/olegrom32/tech-assignment/domain/shared"
)

type GroupRepository interface {
	Find() ([]shared.GroupID, error)
}

type Response struct {
	List []uint64
}

type Handler struct {
	groupRepository GroupRepository
}

func NewHandler(groupRepository GroupRepository) *Handler {
	return &Handler{
		groupRepository: groupRepository,
	}
}

func (h *Handler) Handle() (Response, error) {
	data, err := h.groupRepository.Find()
	if err != nil {
		return Response{}, err
	}

	resp := Response{
		List: make([]uint64, len(data)),
	}

	for k, v := range data {
		resp.List[k] = v.Uint64()
	}

	return resp, nil
}
