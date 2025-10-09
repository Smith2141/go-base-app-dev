package main

import "fmt"

func main() {
	wind := 6         // сила ветра
	rain := false     // дождя нет
	temperature := 16 // температура

	if (!rain || wind <= 4) && temperature > 22 {
		fmt.Println("Идём гулять, на улице хорошо")
	} else {
		fmt.Println("Сидим дома, читаем Практикум")
	}
}
