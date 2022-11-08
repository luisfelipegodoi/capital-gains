package financial

import (
	"github.com/shopspring/decimal"
)

const (
	OperationBuy  = "buy"
	OperationSell = "sell"
)

type Operation struct {
	Type     string          `json:"operation"`
	UnitCost decimal.Decimal `json:"unit-cost"`
	Quantity int64           `json:"quantity"`
	TaxPaid  decimal.Decimal `json:"tax-paid"`
}

type Fee struct {
	Tax decimal.Decimal `json:"tax"`
}
