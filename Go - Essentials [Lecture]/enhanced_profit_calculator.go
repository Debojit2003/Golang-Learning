package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1> Validate user input
//       -> Show error message and exit if invalid input is provided.
//       -> No negative numbers
//       -> Not 0
// 2> Store calculated results into file.
//

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64
	//fmt.Print("Please enter the revenue amount: ")
	//fmt.Scan(&revenue)
	revenue, err1 := getUserInput("Revenue: ")

	// if err != nil {
	// 	fmt.Println((err))
	// 	return
	// }
	//fmt.Print("Please enter the expenses amount: ")
	//fmt.Scan(&expenses)
	expenses, err2 := getUserInput("Expenses: ")

	// if err != nil {
	// 	fmt.Println((err))
	// 	return
	// }

	tax_rate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	//fmt.Print("Please enter the tax rate: ")
	//fmt.Scan(&tax_rate)
	//var ebt float64 = revenue - expenses
	//var profit float64 = ebt * (1 - (tax_rate / 100))
	//var ratio float64 = ebt / profit
	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax_rate)
	fmt.Printf("The EBT is: %.1f\n", ebt)
	fmt.Printf("The profit is: %.1f\n", profit)
	fmt.Printf("The ratio of EBT to profit is: %.3f", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {

	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("Results.txt", []byte(results), 0644)
}

func getUserInput(infotext string) (float64, error) {
	var userinput float64
	fmt.Print(infotext)
	fmt.Scan(&userinput)
	if userinput <= 0 {
		return 0, errors.New("Value Must be a positive number!")
	}

	return userinput, nil
}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (tax_rate / 100))
	ratio := ebt / profit
	return ebt, profit, ratio

}
