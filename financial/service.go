package financial

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"os"
)

const (
	maximumProfitAllowed = 20000
	taxOperationBuy      = 0
)

// CalculateCapitalGains - service method for calculate capital gains
func CalculateCapitalGains(operationList *os.File) ([]Fee, error) {

	operations, err := buildJsonToStruct(operationList)
	if err != nil {
		log.Fatal("Error to convert json file to our struct. Please try again.")
		return nil, err
	}

	fmt.Println("Operations list successfully read")

	fees, err := calculateFees(operations)
	if err != nil {
		log.Fatal("Error to convert fees")
		return nil, err
	}

	fmt.Println("Successfully calculated fees")

	return fees, nil
}

func buildJsonToStruct(operationList *os.File) ([]Operation, error) {
	var operations []Operation

	if err := json.NewDecoder(operationList).Decode(&operations); err != nil {
		return nil, err
	}

	return operations, nil
}

func calculateFees(operations []Operation) ([]Fee, error) {

	var weightedAveragePrice, lastUnitCost, lastDamage decimal.Decimal
	var currentTotalActions int64
	arrFees := make([]Fee, len(operations))

	for i, v := range operations {
		switch v.Type {
		case OperationBuy:
			v.TaxPaid, weightedAveragePrice = calculateBuyTax(v, currentTotalActions, weightedAveragePrice)
			currentTotalActions += v.Quantity
		case OperationSell:
			v.TaxPaid, lastDamage = calculateSellTax(v, currentTotalActions, weightedAveragePrice, lastUnitCost, lastDamage)
			currentTotalActions -= v.Quantity
		}

		lastUnitCost = v.UnitCost
		arrFees[i].Tax = v.TaxPaid
	}

	return arrFees, nil
}

func calculateBuyTax(operation Operation, currentTotalActions int64,
	currentWeightedAverage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {

	newWeightedAverage := calculateWeightedAverage(operation, currentTotalActions, currentWeightedAverage)

	return decimal.NewFromInt(taxOperationBuy), newWeightedAverage
}

func calculateSellTax(operation Operation, currentTotalActions int64,
	weightedAveragePrice, lastUnitCost, lastDamage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {

	loss, taxToPay := calculateTaxOperation(operation, currentTotalActions, weightedAveragePrice,
		lastUnitCost, lastDamage)

	return taxToPay, loss
}

func calculateTaxOperation(operation Operation, currentTotalActions int64, weightedAveragePrice,
	lastUnitCost, lastDamage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {

	var profits, loss, taxToPay decimal.Decimal
	var taxPercentPaid int64 = 20

	actionsDecreased := currentTotalActions - operation.Quantity
	if actionsDecreased == 0 {
		actionsDecreased = currentTotalActions
	}

	if transactionWithProfits := decimal.NewFromInt(operation.UnitCost.IntPart()).GreaterThan(lastUnitCost); transactionWithProfits == true {
		profits = decimal.NewFromInt(actionsDecreased*(operation.UnitCost.IntPart()-lastUnitCost.IntPart()) - lastDamage.IntPart())
		loss = decimal.NewFromInt(0)
		taxToPay = decimal.NewFromInt(profits.IntPart() * taxPercentPaid).Div(decimal.NewFromInt(100))
	} else if decimal.NewFromInt(operation.UnitCost.IntPart()).Equals(weightedAveragePrice) {
		loss = decimal.NewFromInt(0)
		profits = decimal.NewFromInt(0)
	} else {
		loss = decimal.NewFromInt(actionsDecreased * operation.UnitCost.IntPart())
		profits = decimal.NewFromInt(0)
	}

	if isTaxFreeOperation := operation.Quantity * operation.UnitCost.IntPart(); isTaxFreeOperation <= maximumProfitAllowed {
		return loss, decimal.NewFromInt(0)
	}

	return loss, taxToPay
}

func calculateWeightedAverage(operation Operation, currentTotalActions int64,
	currentWeightedAverage decimal.Decimal) decimal.Decimal {

	newWeightedAverage := ((currentTotalActions * currentWeightedAverage.IntPart()) +
		(operation.Quantity * operation.UnitCost.IntPart())) /
		(currentTotalActions + operation.Quantity)

	return decimal.NewFromInt(newWeightedAverage)
}
