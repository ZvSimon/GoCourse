package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	const taxRate float64 = 30

	fmt.Print("Enter revenue : ")
	fmt.Scan(&revenue)

	fmt.Print("\nEnter expenses :")
	fmt.Scan(&expenses)

	earningBeforeTax := revenue - expenses
	earningAfterTax := (revenue - expenses) * (1 - taxRate/100)

	fmt.Print("\nEarning Before Taxe :", earningBeforeTax)
	fmt.Print("\nEarning After Taxe : ", earningAfterTax)

	ratio := earningBeforeTax / earningAfterTax
	fmt.Print("\nRatio is :", ratio)
}
