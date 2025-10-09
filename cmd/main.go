package main

import "fmt"

func main() {
	mark := 4.85

	if 4.7 <= mark && mark <= float64(5) {
		fmt.Println("Отличник")
	}

	if 3.9 <= mark && mark < 4.7 {
		fmt.Println("Ударник")
	}

	if 3.9 <= mark && mark < 4.7 {
		fmt.Println("Троечник")
	}

	if 3.9 <= mark && mark < 4.7 {
		fmt.Println("Тихий троечник")
	}
}
