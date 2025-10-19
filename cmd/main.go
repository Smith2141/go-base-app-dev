package main

import "fmt"

func main() {
	// оценки по отдельным ученикам
	marks := [][]int{
		{5, 4, 5, 5},
		{3, 4, 4, 5, 3},
		{2, 3, 3},
		{5, 5, 4},
		{4, 3, 4, 4, 3},
	}
	var average float32 // итоговый средний балл

	// добавьте код для подсчёта среднего балла
	for _, student := range marks {
		average += calcAvg(student)
	}

	average /= float32(len(marks))

	fmt.Printf("%.2f", average)
}

func calcAvg(markList []int) float32 {
	var (
		markSum int
		result  float32
	)

	for _, mark := range markList {
		markSum += mark
	}

	result = float32(markSum) / float32(len(markList))
	return result
}
