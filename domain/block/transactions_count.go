package block

type TransactionsCount int

func NewTransactionsCount(value int) TransactionsCount {
	return TransactionsCount(value)
}

func (vo TransactionsCount) Int() int {
	return int(vo)
}
