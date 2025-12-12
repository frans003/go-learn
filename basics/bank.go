package main

//--Description--
//Ask the user for what operation he wants to perform at the bank
//Input will be based on a pre defined list with numbers
//User can continue interacting with the bank, even when he has done an action, until they choose exit
//We correct the user in case bad values are inserted, like negative amounts
//--END--

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	accountBalance := getBalanceFromFile()

	fmt.Print("--GO Bank--\n")

	for {
		fmt.Print("What do you want to do?\n\n")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("You choosed: ", choice)

		switch choice {
		case 1:
			fmt.Printf("\nBalance: €%.2f\n\n", accountBalance)
		case 2:
			var deposit float64
			fmt.Print("How much do you want to deposit? ")
			fmt.Scan(&deposit)
			if deposit > 0 {
				accountBalance += deposit
				fmt.Printf("New balance: €%.2f\n\n", accountBalance)
				writeBalanceToFile(accountBalance)
			} else {
				fmt.Println("Invalid amount")
				continue
			}
		case 3:
			var withdrawal float64
			fmt.Print("How much do you want to withdraw? ")
			fmt.Scan(&withdrawal)
			if withdrawal > accountBalance {
				fmt.Println("Balance too low")
			} else if withdrawal < accountBalance && withdrawal > 0 {
				accountBalance -= withdrawal
				writeBalanceToFile(accountBalance)
				fmt.Printf("New balance: €%.2f", accountBalance)
			} else {
				fmt.Println("Invalid option, try again")
				continue
			}
		default:
			fmt.Print("Thanks for choosing Go Bank")
			return
		}
	}
}
