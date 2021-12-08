package index

import (
	"github.com/olegrom32/tech-assignment/domain/shared"
)

type Model struct {
	ID                shared.IndexID
	Name              Name
	ETHPriceInWei     ETHPrice
	USDPriceInCents   USDPrice
	USDCapitalization USDCapitalization
	PercentageChange  PercentageChange
}
