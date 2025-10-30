package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Получаем пользовательский ввод
func getData() string {
	var secondInput string
	fmt.Print("Введите оператор и второй операнд, или \"quit\": ")
	fmt.Scan(&secondInput)

	return secondInput
}

// Выполняем парсинг строки и вычисления
func calculate(firstDigit int, data string) int {
	var result int
	var runes []rune = []rune(data)

	if len(runes) < 2 {
		panic("Проверьте правильность ввода!!!")
	}

	secondDigit, err := strconv.Atoi(strings.TrimPrefix(data, string(runes[0])))

	if err != nil {
		panic(err)
	}

	var operator string = strings.TrimSuffix(data, fmt.Sprint(secondDigit))

	switch operator {
	case "+":
		result = firstDigit + secondDigit
	case "-":
		result = firstDigit - secondDigit
	case "*":
		result = firstDigit * secondDigit
	case "/":
		result = firstDigit / secondDigit
	}

	return result
}

func main() {
	var firstDigit int
	var secondData string

	fmt.Print("Введите первый операнд: ")
	fmt.Scan(&firstDigit)

	secondData = getData()

	for secondData != "q" {
		fmt.Println("Выполнение программы")

		firstDigit = calculate(firstDigit, secondData)
		fmt.Println(firstDigit)
		secondData = getData()
	}

	fmt.Println("Выполняю выход из программы")
}
