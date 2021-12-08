package shared

type IndexID uint64

func NewIndexID(value int64) IndexID {
	return IndexID(value)
}

func (vo IndexID) Uint64() uint64 {
	return uint64(vo)
}

func (vo IndexID) Int64() int64 {
	return int64(vo)
}
