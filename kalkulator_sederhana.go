package main

import "fmt"

func main() {

	var number1, number2 float64
	var operator string

	fmt.Print("Enter the first number : ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the operator ( + - * / ) : ")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Print("%.3f %s %.3f = %.3f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by Zero Situation")
		} else {
			fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1/number2)
		}

	default:
		fmt.Println("Invalid Operator")

	}
}
