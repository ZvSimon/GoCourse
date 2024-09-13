package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var annualInterestRate = 5.5
	var years float64 = 10

	var futureInvestmentValue = investmentAmount * math.Pow(1+annualInterestRate/100, years)
	fmt.Print("Future value is ", futureInvestmentValue)

}
