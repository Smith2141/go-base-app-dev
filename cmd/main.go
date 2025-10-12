package main

import "fmt"

const (
	Limit13 = 5000000
)

func main() {
	// income := 7878415.0 // доход
	income := 5_500_000.0 // доход
	tax := 0.0            // налог

	// вставьте код, который вычисляет налог
	var part15 float64 = income - float64(Limit13)
	if part15 >= 0 {
		tax = (part15 / 100) * 15

		var part13 float64 = income - part15
		tax += (part13 / 100) * 13
	} else {
		tax = (income / 100) * 13
	}

	fmt.Printf("Доход: %.2f, НДФЛ: %.2f", income, tax)
}
