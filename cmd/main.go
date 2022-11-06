package main

import (
	"fmt"
	"github.com/luisfelipegodoi/capital-gains/financial"
	"log"
	"os"
)

func main() {

	// Receiving operations list
	fmt.Println("Please input the operations list:")
	operationList := os.Stdin
	
	// Calculate capital gains
	capitalGains, err := financial.CalculateCapitalGains(operationList)
	if err != nil {
		log.Fatal("Error to calculate capital gains")
	}

	// Printing result
	fmt.Println("Result:")
	fmt.Printf("%+v\n", capitalGains)
}
