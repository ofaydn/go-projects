package main

import (
	"errors"
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot make operation with 0")
	} else {
		return a / b, nil
	}

}

func main() {

	var a, b float64
	var operation string
	var result float64
	var err error
	fmt.Println("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&b)
	fmt.Println("Enter operation: ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		result = add(a, b)
	case "-":
		result = sub(a, b)
	case "*":
		result = mul(a, b)
	case "/":
		result, err = div(a, b)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Printf("Result: %f\n", result)

}
