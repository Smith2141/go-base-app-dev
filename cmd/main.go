package main

import "fmt"

func main() {
	var result string

	for i := 1; i <= 100; i++ {
		result = resultModify(i)

		fmt.Println(result)
	}
}

func resultModify(digit int) string {

	var resultElem string

	if digit%3 == 0 {
		resultElem += "Fizz"
	}
	if digit%5 == 0 {
		resultElem += "Buzz"
	}

	if digit%3 != 0 && digit%5 != 0 {

		resultElem = fmt.Sprint(digit)
	}

	return resultElem
}
