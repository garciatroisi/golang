package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
 
	revenue = getInput("Revenue:")
	expenses = getInput("Expenses:")
	taxRate = getInput("Tax Rate:")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)
	fmt.Printf("EBT:%.1f\n", ebt)
	fmt.Printf("Profit:%.1f\n", profit)
	fmt.Printf("Ratio:%.3f\n", ratio)
}


func getInput(msj string) (float64) {
	var userInput float64
	fmt.Print(msj)
	fmt.Scan(&userInput)
	return userInput
}

func calculate(revenue , expenses , taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}