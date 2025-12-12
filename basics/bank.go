package main

import "fmt"

func main() {
	//--Description--
	//Ask the user for what operation he wants to perform at the bank
	//Input will be based on a pre defined list with numbers
	//User can continue interacting with the bank, even when he has done an action, until they choose exit
	//We correct the user in case bad values are inserted, like negative amounts
	//--END--

	accountBalance := 1000.0

	fmt.Println("--GO Bank--\n")

	for {
		fmt.Println("What do you want to do?\n")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("You choosed: ", choice)

		if choice == 1 {
			fmt.Printf("Your balance is: €%.2f", accountBalance)
		} else if choice == 2 {
			var deposit float64
			fmt.Print("How much do you want to deposit? ")
			fmt.Scan(&deposit)
			accountBalance += deposit
			fmt.Printf("New balance: €%.2f\n\n", accountBalance)
		} else if choice == 3 {
			var withdrawal float64
			fmt.Print("How much do you want to withdraw? ")
			fmt.Scan(&withdrawal)
			if withdrawal > accountBalance {
				fmt.Println("Balance too low")
			} else if withdrawal < accountBalance && withdrawal > 0 {
				accountBalance -= withdrawal
				fmt.Printf("New balance: €%.2f", accountBalance)
			} else {
				fmt.Println("Invalid option, try again")
				return
			}
		} else if choice == 4 {
			fmt.Print("Goodbye")
			break
		} else {
			fmt.Println("Invalid option, try again")
		}
	}
	fmt.Print("Thanks for choosing Go Bank")

}
