package index

type PercentageChange int64

func NewPercentageChange(value int64) PercentageChange {
	return PercentageChange(value)
}

func (vo PercentageChange) Int64() int64 {
	return int64(vo)
}
