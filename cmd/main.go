package main

import "fmt"

const (
	numberOfCubeSides int = 6
	numberOfCubes     int = 8
)

// функция для вычисления площади куба
func calcCubeArea(side int) int {
	// формулу для вычисления площади одной грани куба Афанасий написал:
	oneFace := side * side

	// вычислите и возвратите полную площадь куба: у него шесть одинаковых граней
	cubeArea := oneFace * numberOfCubeSides

	return cubeArea
}

func main() {
	// присвойте переменной oneCubeArea значение,
	// которое вернёт функция calcCubeArea() с аргументом 3:
	// 3 метра — это длина ребра куба
	oneCubeArea := calcCubeArea(3)

	// вычислите общую площадь стекла, необходимого
	// для строительства 8 кубов,
	// и сохраните это значение в переменную fullArea
	fullArea := oneCubeArea * numberOfCubes

	fmt.Println("Необходимая площадь стекла для 8 кубов, кв. м:", fullArea)
}
