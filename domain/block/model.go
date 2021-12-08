package block

import (
	"time"
)

type Model struct {
	Number            Number
	Hash              Hash
	GasUsed           GasUsed
	ReceivedAt        time.Time
	TransactionsCount TransactionsCount
}
