// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	shared "github.com/olegrom32/tech-assignment/domain/shared"
	mock "github.com/stretchr/testify/mock"
)

// GroupRepository is an autogenerated mock type for the GroupRepository type
type GroupRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields:
func (_m *GroupRepository) Find() ([]shared.GroupID, error) {
	ret := _m.Called()

	var r0 []shared.GroupID
	if rf, ok := ret.Get(0).(func() []shared.GroupID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]shared.GroupID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
