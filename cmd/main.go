package main

import "fmt"

func main() {
	// Объявляем целочисленную переменную
	var number int
	// Объявляем две строковые переменные
	var hello, bye string

	// Присваиваем переменной number число 4
	number = 4

	// Присваиваем переменной hello строку "Привет!"
	hello = "Привет!"
	// Присваиваем переменной bye строку "Пока!"
	bye = "Пока!"

	// Выводим значения перменных на экран
	fmt.Println("number = ", number)
	fmt.Println("hello = ", hello)
	fmt.Println("bye = ", bye)
}
