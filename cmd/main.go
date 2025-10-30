package main

import "fmt"

// Получаем и разбираем пользовательский ввод
func getData() (string, float64) {
	var operator string
	var secondDigit float64

	fmt.Print("Введите оператор: ")
	fmt.Scan(&operator)
	fmt.Print("Введите второй операнд: ")
	fmt.Scan(&secondDigit)

	return operator, secondDigit
}

func closureCalculate() func() float64 {
	var firstDigit float64
	var firstDigitEntered bool

	return func() float64 {

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
			if secondDigit == 0.0 {
				fmt.Println("На ноль делить нельзя")
				return 0.0
			}
			firstDigit = firstDigit / secondDigit
		default:
			fmt.Println("Оператор не корректен")
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
