package listgroups

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/olegrom32/tech-assignment/application"
	"github.com/olegrom32/tech-assignment/application/handler/listgroups/mocks"
	"github.com/olegrom32/tech-assignment/domain/shared"
)

//go:generate ../../../bin/mockery --all

func TestHandler_HandleSuccess(t *testing.T) {
	r := &mocks.GroupRepository{}
	h := NewHandler(r)

	r.On("Find").Return([]shared.GroupID{1, 3, 5}, nil)

	res, err := h.Handle()
	require.NoError(t, err)

	assert.Len(t, res.List, 3)
	assert.Equal(t, uint64(1), res.List[0])
	assert.Equal(t, uint64(3), res.List[1])
	assert.Equal(t, uint64(5), res.List[2])
}

func TestHandler_HandleFail(t *testing.T) {
	r := &mocks.GroupRepository{}
	h := NewHandler(r)

	r.On("Find").Return(nil, application.ErrNoGroupsFound)

	_, err := h.Handle()

	assert.ErrorIs(t, err, application.ErrNoGroupsFound)
}
