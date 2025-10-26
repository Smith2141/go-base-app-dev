package main

import "fmt"

const (
	cubeSideCount int = 12
	numberOfCubes int = 8
)

// функция для вычисления периметра куба
func calcCubePerimeter(side int) int {
	return side * cubeSideCount
}

func main() {
	// присвойте переменной oneCubePerimeter значение,
	// которое вернёт функция calcCubePerimeter() с аргументом 3:
	// 3 метра — это длина ребра куба

	oneCubePerimeter := calcCubePerimeter(3)

	// вычислите общую длину палок, необходимых
	// для строительства 8 кубов,
	// и сохраните это значение в переменную fullLength
	fullLength := numberOfCubes * oneCubePerimeter

	// теперь напечатаем результат (в этой строке ничего изменять не нужно)
	fmt.Println("Необходимый метраж палок для 8 кубов:", fullLength)
}
