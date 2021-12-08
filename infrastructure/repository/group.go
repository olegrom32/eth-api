package repository

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"

	"github.com/olegrom32/tech-assignment/application"
	"github.com/olegrom32/tech-assignment/domain/group"
	"github.com/olegrom32/tech-assignment/domain/shared"
	"github.com/olegrom32/tech-assignment/pkg/contract"
)

type Group struct {
	client *contract.ContractCaller
}

func NewGroup(client *contract.ContractCaller) *Group {
	return &Group{
		client: client,
	}
}

func (r *Group) Find() ([]shared.GroupID, error) {
	ids, err := r.client.GetGroupIds(&bind.CallOpts{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	n := len(ids)
	if n == 0 {
		return nil, errors.WithStack(application.ErrNoGroupsFound)
	}

	res := make([]shared.GroupID, n)

	for k, v := range ids {
		res[k] = shared.NewGroupID(v.Int64())
	}

	return res, nil
}

func (r *Group) FindByID(groupID shared.GroupID) (group.Model, error) {
	data, err := r.client.GetGroup(&bind.CallOpts{}, big.NewInt(groupID.Int64()))
	if err != nil {
		return group.Model{}, errors.WithStack(err)
	}

	m := group.Model{}

	if m.Name, err = group.NewName(data.Name); err != nil {
		return group.Model{}, errors.WithStack(err)
	}

	m.Indexes = make([]shared.IndexID, len(data.Indexes))

	for k, v := range data.Indexes {
		m.Indexes[k] = shared.NewIndexID(v.Int64())
	}

	return m, nil
}
