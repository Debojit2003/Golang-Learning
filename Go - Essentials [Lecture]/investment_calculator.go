package main

import (
	"fmt"
	"math"
)

const inflation_rate float64 = 2.5

func main() {
	var investmentAmount float64
	//var investmentAmount, years float64 = 1000, 10
	//expectedReturnRate := 5.5
	var expectedReturnRate float64
	//var years float64 = 10
	var years float64

	fmt.Print("Enter an investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFuturevalues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//var futureRealValue = futureValue / math.Pow(1+inflation_rate/100, years)

	var formattedFV = fmt.Sprintf("The future value is: %.2f\n", futureValue)
	var formattedFRV = fmt.Sprintf("The future real value is: %.2f\n", futureRealValue)

	//fmt.Printf("The future value is: %.2f\n", futureValue)
	//fmt.Println("The future real value is: ", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
	//fmt.Printf(`The future value is: %.2f\n
	//The future real value is: %.2f`, futureValue, futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFuturevalues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflation_rate/100, years)
	// return fv, rfv
	return //Go automatically returns the elements fv and rfv
}
