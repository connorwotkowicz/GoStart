package main

import (
	"fmt"
)

func main() {
	num1 := getFloat("Enter first number: ")
	operator := getOperator("Enter an operator (+, -, *, /): ")
	num2 := getFloat("Enter second number: ")

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}

func getFloat(prompt string) float64 {
	var num float64
	fmt.Print(prompt)
	fmt.Scanln(&num)
	return num
}

func getOperator(prompt string) string {
	var op string
	fmt.Print(prompt)
	fmt.Scanln(&op)
	return op
}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}
