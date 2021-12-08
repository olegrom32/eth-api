package block

type Number uint64

func NewNumber(value uint64) Number {
	return Number(value)
}

func (vo Number) Uint64() uint64 {
	return uint64(vo)
}

func (vo Number) Int64() int64 {
	return int64(vo)
}
