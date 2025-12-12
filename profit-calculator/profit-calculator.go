package main

import "fmt"

func main() {

	var revenues, expenses, taxRate float64 = calcProfit()

	EBT := calcEarningsBeforeTaxes(revenues, expenses)
	profits, taxes := calcProfits(EBT, taxRate)
	ratio := calcRatio(EBT, profits)
	fmt.Printf(`Your revenues before taxes are: €%.2f
Total taxes: €%.2f
Earnings after tax: €%.2f
Ratio: %.1f`, EBT, taxes, profits, ratio)

	formattedProfit := fmt.Sprintf("\nRevenues: €%.2f", profits)
	fmt.Print(formattedProfit)
}

func calcProfit() (revenues, expenses, taxRate float64) {
	fmt.Print("What are the company revenues? ")
	fmt.Scan(&revenues)

	fmt.Print("What are the total expenses? ")
	fmt.Scan(&expenses)

	fmt.Print("What is the tax rate? ")
	fmt.Scan(&taxRate)

	return revenues, expenses, taxRate
}

func calcEarningsBeforeTaxes(revenues, expenses float64) (ebt float64) {
	ebt = revenues - expenses
	return ebt
}

func calcProfits(EBT, taxRate float64) (profits, taxes float64) {
	taxes = EBT * (taxRate / 100)
	profits = EBT - taxes
	return profits, taxes
}

func calcRatio(EBT, profits float64) (ratio float64) {
	ratio = EBT / profits
	return ratio
}
