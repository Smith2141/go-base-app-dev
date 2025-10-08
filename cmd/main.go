package main

import (
	"fmt"
	"strconv" // импортируем пакет, в котором находится функция Atoi()
)

func main() {
	// это строковая переменная
	hello := "Привет"

	someInt, err := strconv.Atoi(hello)
	// Если ошибка не равна `nil`, то вызовем панику. Программа аварийно завершится.
	if err != nil {
		panic(err)
	}

	fmt.Println(someInt)
}
