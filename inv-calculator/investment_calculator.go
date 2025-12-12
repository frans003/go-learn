package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	const inflationRate = 2.5

	var investedAmount, years, expectedReturnRate float64

	fmt.Print("How much do you want to invest? ")
	fmt.Scan(&investedAmount)

	fmt.Print("How many years? ")
	fmt.Scan(&years)

	fmt.Print("What is the expected return rate (in %)? ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investedAmount, expectedReturnRate, years)

	fmt.Print("Future Value: ")
	fmt.Println(math.Round(futureValue))
	fmt.Print("Future Real Value: ")
	fmt.Println(math.Round(futureRealValue))
	result := sum(1, 1)
	fmt.Println(result)
}

func sum(a int, b int) int {
	return a + b
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
