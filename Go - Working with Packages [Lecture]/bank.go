package main

//go get github.com/Pallinder/go-randomdata

import (
	"bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------------")
		// panic("Can't continue, due to errors!")// return //Directly gets out if error occurs!

	}
	fmt.Println("Welcome to GO Bank!")
	fmt.Println("Reach us 24/7 at ", randomdata.PhoneNumber()) //Using package from external the link is provided above.
	for {

		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 4:
			fmt.Println("Thank you! Have a nice day.")
			fmt.Println("Thank you for choosing our bank!🔥")
			// break
			return

		}

	}

}
