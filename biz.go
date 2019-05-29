package main

import "fmt"

func main() {

	 var Expenses float64;
	 Expenses = 9500

	fmt.Print(" £ ",breakEvenPoint(Expenses,salePriceFromGrossProfit(0.5,1),1))
	fmt.Print(" -")
	fmt.Print(" £ ",breakEvenPoint(Expenses,salePriceFromGrossProfit(0.7,1),1))
}

func breakEvenPoint(FixedCosts, SalePrice, VariableCost float64) int {

	var ContributionMargin = SalePrice - VariableCost
	var BreakEvenPoint = FixedCosts / ContributionMargin

	return int(BreakEvenPoint)
}

func grossProfitMargin(Revenue, CostOfSale float64) float64 {

	var GrossProfitMargin = (Revenue - CostOfSale) / Revenue

	return float64(GrossProfitMargin)
}

func salePriceFromGrossProfit(GrossProfitMargin, CostOfSale float64) float64 {
	var SalePrice = CostOfSale + (CostOfSale * GrossProfitMargin)

	return float64(SalePrice)
}

