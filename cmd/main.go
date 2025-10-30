package main

import (
	"fmt"
)

// Функция, которая принимает другую функцию в качестве аргумента
func squareAndPrint(fn func(int) int) {
	result := fn(5)
	fmt.Println(result)
}

// Функция, которая будет передаваться как аргумент
func square(x int) int {
	return x * x
}

func main() {
	// Передаем функцию square в doSomething
	squareAndPrint(square)
}
