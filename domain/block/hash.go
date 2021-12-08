package block

type Hash string

func NewHash(value string) Hash {
	return Hash(value)
}

func (vo Hash) String() string {
	return string(vo)
}
