package main

import "fmt"

func main() {
	var firstPart string = "У вас "
	var secondPart string = " новых сообщений"

	var count int = 8
	var countString string = fmt.Sprint(count)

	fmt.Println(firstPart + countString + secondPart)
}
