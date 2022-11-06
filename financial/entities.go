package financial

import (
	"github.com/shopspring/decimal"
)

type Operation struct {
	Type     string          `json:"operation"`
	UnitCost decimal.Decimal `json:"unit-cost"`
	Quantity int32           `json:"quantity"`
}

type Fee struct {
	Tax decimal.Decimal `json:"tax"`
}
