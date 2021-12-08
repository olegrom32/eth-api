package block

type GasUsed uint64

func NewGasUsed(value uint64) GasUsed {
	return GasUsed(value)
}

func (vo GasUsed) Uint64() uint64 {
	return uint64(vo)
}
