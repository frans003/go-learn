package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investedAmount, years, expectedReturnRate float64

	fmt.Print("How much do you want to invest? ")
	fmt.Scan(&investedAmount)

	fmt.Print("How many years? ")
	fmt.Scan(&years)

	fmt.Print("What is the expected return rate (in %)? ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investedAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

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
