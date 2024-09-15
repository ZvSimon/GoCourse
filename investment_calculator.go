package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureInvestmentValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureInvestmentValue / math.Pow(1+inflationRate/100, years)
	fmt.Print("Future value is ", futureInvestmentValue)
	fmt.Print("\nFuture real value is ", futureRealValue)

}
