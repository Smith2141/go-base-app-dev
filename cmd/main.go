package main

import "fmt"

const (
	numberOfCubes int = 2
	cubeSideValue int = 3
)

// функция для вычисления периметра куба
func calcCubePerimeter(side int) int {
	return side * 12
}

// функция для вычисления площади куба
func calcCubeArea(side int) int {
	oneFace := side * side
	return oneFace * 6
}

// дополните объявление функции:
// теперь она должна принимать два параметра —
// длину ребра куба и количество кубов
func calcCube(side int, amount int) {
	// вызываем функцию, рассчитывающую периметр,
	// и передаём в неё размер куба
	oneCubePerimeter := calcCubePerimeter(side)
	// здесь вместо многоточия должна стоять переменная,
	// хранящая количество кубов, переданное во втором аргументе
	fullLength := oneCubePerimeter * amount

	// вызываем функцию, рассчитывающую площадь стекла,
	// и передаём в неё размер куба
	oneCubeArea := calcCubeArea(side)
	// здесь вместо многоточия должна стоять переменная,
	// хранящая количество кубов, переданное во втором аргументе
	fullArea := oneCubeArea * amount

	// поставьте вместо многоточий переменную с количеством кубов
	// и переменную с длиной ребра
	fmt.Printf("Для %d кубов с ребром %d м понадобится"+
		" %d м палок и %d кв. м стекла",
		amount, side, fullLength, fullArea)
}

func main() {
	calcCube(cubeSideValue, numberOfCubes)
}
