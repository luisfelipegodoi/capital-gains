package financial

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

	arrFees := make([]Fee, len(operations))

	for i, v := range operations {
		arrFees[i].Tax = v.UnitCost
	}

	return arrFees, nil
}
