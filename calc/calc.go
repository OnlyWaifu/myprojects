package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operator string

	fmt.Println("Введите первое число: ")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе число: ")
	fmt.Scanln(&num2)

	fmt.Println("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Нельзя делить на ноль")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f", num1, num2, num1/num2)
		}
	default:
		fmt.Println("Неверный оператор")
	}
}
