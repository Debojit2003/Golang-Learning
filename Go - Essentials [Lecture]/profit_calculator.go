package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64
	//fmt.Print("Please enter the revenue amount: ")
	//fmt.Scan(&revenue)
	revenue = getUserInput("Revenue: ")
	//fmt.Print("Please enter the expenses amount: ")
	//fmt.Scan(&expenses)
	expenses = getUserInput("Expenses: ")
	tax_rate = getUserInput("Tax Rate: ")
	//fmt.Print("Please enter the tax rate: ")
	//fmt.Scan(&tax_rate)
	//var ebt float64 = revenue - expenses
	//var profit float64 = ebt * (1 - (tax_rate / 100))
	//var ratio float64 = ebt / profit
	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax_rate)
	fmt.Println("The EBT is: ", ebt)
	fmt.Println("The profit is: ", profit)
	fmt.Printf("The ratio of EBT to profit is: %.1f", ratio)
}

func getUserInput(infotext string) float64 {
	var userinput float64
	fmt.Print(infotext)
	fmt.Scan(&userinput)
	return userinput
}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (tax_rate / 100))
	ratio := ebt / profit
	return ebt, profit, ratio

}
