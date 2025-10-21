package main

import "fmt"

func main() {
	towns := map[string]int{
		"Омск":      1104485,
		"Киров":     475464,
		"Астрахань": 465524,
		"Пятигорск": 144955,
		"Пенза":     488299,
		"Иркутск":   606369,
	}
	towns2 := map[string]int{
		"Якутск":    384979,
		"Омск":      1104485,
		"Хабаровск": 615570,
		"Челябинск": 1177058,
		"Астрахань": 465524,
		"Пенза":     488299,
	}

	for name, people := range towns2 {
		_, ok := towns[name]
		if !ok {
			towns[name] = people
		}
	}
	fmt.Println(towns)
}
