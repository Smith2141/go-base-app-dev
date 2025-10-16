package main

import "fmt"

func main() {
	var hidden string // результирующая строка со звёздочками
	var isDogAhead bool = true
	email := `vasyapupkin33@mail.ru`

	for i, ch := range email {
		if string(ch) == "@" {
			isDogAhead = false
		}

		if isDogAhead && i > 1 {
			hidden += "*"
		} else {
			hidden += string(ch)
		}
	}

	fmt.Println(hidden)
}
