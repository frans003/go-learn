package main

import "fmt"

func main() {

	var revenues, expenses, taxRate float64

	fmt.Print("What are the company revenues? ")
	fmt.Scan(&revenues)
	fmt.Print("What are the total expenses? ")
	fmt.Scan(&expenses)
	fmt.Print("What is the tax rate? ")
	fmt.Scan(&taxRate)

	taxRate = (taxRate / 100)

	EBT := revenues - expenses
	profit := EBT - (EBT * taxRate)
	ratio := EBT / profit

	fmt.Print("Your revenues before taxes are: ")
	fmt.Println(EBT)

	fmt.Print("Earning after tax: ")
	fmt.Println(profit)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}
