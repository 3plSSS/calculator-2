package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Простой калькулятор")

	var num1Str, num2Str string
	var operation string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1Str)
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		panic("Ошибка: введите корректное число")
	}

	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scanln(&operation)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2Str)
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		panic("Ошибка: введите корректное число")
	}

	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("Ошибка: деление на ноль невозможно")
		}
		result = num1 / num2
	default:
		panic(fmt.Sprintf("Ошибка: неизвестная операция '%s'", operation))
	}

	fmt.Printf("Результат: %.2f\n", result)
}