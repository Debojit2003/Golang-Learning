//Program for Control Statements

package main

import "fmt"

func main() {

	var accountBalance float64 = 10000
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

		if choice == 1 {
			fmt.Println("Your current balance is: ", accountBalance)
		} else if choice == 2 {
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
		} else if choice == 3 {
			var money float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&money)
			if money <= 0 || money > accountBalance {
				fmt.Println("Invalid amount. Withdrawable amount should be greater than 0 and lesser than the current balance.")
				continue
			}
			accountBalance = accountBalance - money
			fmt.Println("Your current balance after withdraw: ", accountBalance)
		} else {
			fmt.Println("Thank you! Have a nice day.")
			break
		}

	}

	fmt.Println("Thank you for choosing our bank!🔥")

}
