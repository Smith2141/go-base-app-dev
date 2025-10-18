package main

import "fmt"

func main() {
	s := "Мезчхе%зцчхкьексцє%з%9%ьеце%ме%ширус%йусе"
	// расшифруйте сообщение и выведете его
	runes := []rune(s) // преобразуем в слайс рун
	for rune := range runes {
		runes[rune] -= 5
		// fmt.Print(string(runes[rune]))
	}
	message := string(runes)
	fmt.Println(message)

}
