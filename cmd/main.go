package main

import "fmt"

func main() {
	// Первый множитель
	for first := 3; first <= 9; first += 2 {
		// Второй множитель
		for second := 3; second <= 9; second += 2 {
			fmt.Printf("%dx%d = %d\n", first, second, first*second)
		}
	}
}
