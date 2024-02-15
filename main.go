package main

import (
	"fmt"
	"my_calculator_app/calculator"
	"my_calculator_app/utils"
)

func main() {
	fmt.Println("Hello from my-app !")
	resultAdd := calculator.Add(5, 3)
	resultSub := calculator.Subtract(8, 4)
	fmt.Printf("Addition result: %d\n", resultAdd)
	fmt.Printf("Addition result: %d\n", resultSub)
	isValid := utils.Validate("Some Input")
	fmt.Print(isValid)
}
