package main

import "fmt"

func main() {

	fmt.Println(breakEvenPoint(5000,700,300))
}

func breakEvenPoint(fixedCosts, salePrice, variableCost int) int {

	var ContributionMargin = salePrice - variableCost
	var BreakEvenPoint = fixedCosts / ContributionMargin

	return int(BreakEvenPoint)
}