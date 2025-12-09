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

	EBT := revenues - expenses
	profit := EBT - (EBT * (taxRate / 100))
	ratio := EBT / profit

	fmt.Printf(`Your revenues before taxes are: €%.2f
Earnings after tax: €%.2f
Ratio: %.1f`, EBT, profit, ratio)

	formattedProfit := fmt.Sprintf("\nRevenues: €%.2f", profit)
	fmt.Print(formattedProfit)
}
