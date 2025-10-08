package main

import (
	"fmt"
)

func main() {
	user := "молодой падаван"
	lines := 50

	// объявляем строковую переменную и присваиваем ей результат преобразования
	// с помощью функции strconv.Itoa()
	// linesStr := strconv.Itoa(lines)
	linesStr := fmt.Sprint(lines)

	// конкатенируем строки и выводим на экран
	fmt.Println("Поздравляю, " + user + "! Ты написал " + linesStr + " строк кода.")
}
