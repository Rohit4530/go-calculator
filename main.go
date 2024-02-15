package main

import (
	"fmt"

	"github.com/Rohit4530/go-calculator/calculator"
	"github.com/Rohit4530/go-calculator/utils"
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
