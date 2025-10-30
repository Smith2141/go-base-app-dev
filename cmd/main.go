package main

import "fmt"

// функция для вычисления периметра куба
func calcCubePerimeter(side float64) float64 {
	return side * 12
}

// функция для вычисления площади куба
func calcCubeArea(side float64) float64 {
	oneFace := side * side
	return oneFace * 6
}

func calcCube(side float64, amount int) {
	oneCubePerimeter := calcCubePerimeter(side)
	fullLength := oneCubePerimeter * float64(amount)
	oneCubeArea := calcCubeArea(side)
	fullArea := oneCubeArea * float64(amount)

	fmt.Printf("Для %d кубов с ребром %.2f м понадобится"+
		" %.2f м палок и %.2f кв. м стекла\r\n",
		amount, side, fullLength, fullArea)
}

func main() {
	// ниже напишите три вызова функции calcСube(),
	// каждый вызов должен быть на отдельной строке

	calcCube(2, 4)
	calcCube(0.5, 26)
	calcCube(0.61, 6)
}
