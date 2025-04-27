package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 10000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------------")
		// panic("Can't continue, due to errors!")// return //Directly gets out if error occurs!

	}
	fmt.Println("Welcome to GO Bank!")
	for {

		fmt.Println("What are the operations do you want to perform?")
		fmt.Println("1. Check Account Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Please enter your choice to proceed: ")
		var choice int
		fmt.Scan(&choice)
		// fmt.Print("Your entered choice is: ", choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your current balance is: ", accountBalance)

		case 2:
			var money float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&money)
			if money <= 0 {
				fmt.Println("Invalid Amount. Deposit value should be greater than 0")
				//return
				continue
			}
			accountBalance = accountBalance + money
			fmt.Println("Your current balance after deposit: ", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			var money float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&money)
			if money <= 0 || money > accountBalance {
				fmt.Println("Invalid amount. Withdrawable amount should be greater than 0 and lesser than the current balance.")
				continue
			}
			accountBalance = accountBalance - money
			fmt.Println("Your current balance after withdraw: ", accountBalance)
			writeBalanceToFile(accountBalance)

		case 4:
			fmt.Println("Thank you! Have a nice day.")
			fmt.Println("Thank you for choosing our bank!🔥")
			// break
			return

		}

	}

}
