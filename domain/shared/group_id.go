package shared

type GroupID uint64

func NewGroupID(value int64) GroupID {
	return GroupID(value)
}

func (vo GroupID) Uint64() uint64 {
	return uint64(vo)
}

func (vo GroupID) Int64() int64 {
	return int64(vo)
}
