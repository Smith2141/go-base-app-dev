package main

import "fmt"

func main() {
	i := 100
	j := 7
	op := '/'

	var result int
	switch op {
	case '+':
		result = i + j
		// добавьте вычисление разности, произведения и частного
	case '-':
		result = i - j
	case '*':
		result = i * j
	case '/':
		result = i / j
	}
	fmt.Println(i, string(op), j, "=", result)
}
