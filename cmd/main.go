package main

import "fmt"

// printGuests выводит заголовок и список гостей
// определите параметры title и list
func printGuests(title string, list ...string) {
	// напишите тело функции
	fmt.Println(title)
	for num, guest := range list {
		fmt.Println(num+1, guest)
	}
}

func main() {
	printGuests("Список гостей на 4 августа", "Вера", "Ринат",
		"Владимир", "Оксана", "Игорь")
}
