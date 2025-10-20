package main

import "fmt"

const (
	fixSymbol  = '?'
	cLetterEng = 'c'
)

func main() {
	words := []string{"сервер", "cистема", "специалист", "слайc", "процессор",
		"масcив", "строка", "максимум", "cпоcоб", "парсер", "условие"}
	var mistakes []string

	// вставьте недостающий код
	for _, word := range words {
		var hasMistake bool
		runes := []rune(word)
		for idx := range runes {
			if runes[idx] == cLetterEng {
				runes[idx] = fixSymbol
				hasMistake = true
				// fmt.Println("fixed: ", string(runes))
			}
		}

		if hasMistake {
			mistakes = append(mistakes, string(runes))
		}
	}

	fmt.Println(mistakes)
}
