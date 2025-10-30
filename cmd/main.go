package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Получаем и разбираем пользовательский ввод
func getData() (string, int) {
	var secondInput, operator string
	var secondDigit int
	var err error

	fmt.Print("Введите оператор и второй операнд: ")
	fmt.Scan(&secondInput)

	var runes []rune = []rune(secondInput)

	secondDigit, err = strconv.Atoi(strings.TrimPrefix(secondInput, string(runes[0])))

	if err != nil {
		panic("Проверьте правильность ввода!!!")
	}

	operator = strings.TrimSuffix(secondInput, fmt.Sprint(secondDigit))

	return operator, secondDigit
}

func closureCalculate() func() int {
	var firstDigit int
	var firstDigitEntered bool

	return func() int {

		if !firstDigitEntered {
			fmt.Print("Введите первый операнд: ")
			fmt.Scan(&firstDigit)
			firstDigitEntered = true
		}

		operator, secondDigit := getData()

		switch operator {
		case "+":
			firstDigit = firstDigit + secondDigit
		case "-":
			firstDigit = firstDigit - secondDigit
		case "*":
			firstDigit = firstDigit * secondDigit
		case "/":
			firstDigit = firstDigit / secondDigit
		}

		return firstDigit
	}
}

func main() {
	process := closureCalculate()

	for {
		result := process()

		fmt.Println(result)
	}
}
