package main

import "fmt"

func main() {
	age := 50

	var birthday string

	// напишите switch в соответствии с заданием
	switch {
	case age%25 == 0:
		birthday = "С юбилеем!"
	case age%5 == 0:
		birthday = "Поздравляем! У вас круглая дата."
	default:
		birthday = "С днём рождения!"
	}

	fmt.Println(birthday)
}
