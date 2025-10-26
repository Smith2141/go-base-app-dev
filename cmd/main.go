package main

import "fmt"

// []int указывает, что temperatures — это слайс с элементами типа int
func comfortCount(temperatures []int) {
	// напишите код функции
	var count int
	for _, temp := range temperatures {
		if temp >= 22 && temp <= 26 {
			count++
		}
	}

	fmt.Printf("Количество тёплых дней в этом месяце: %d\n", count)
}

func main() {
	// здесь код не меняйте
	may2017 := []int{24, 26, 15, 10, 15, 19, 10, 1, 4, 7, 7, 7, 12, 14, 17,
		8, 9, 19, 21, 22, 11, 15, 19, 23, 15, 21, 16, 13, 25, 17, 19}
	may2018 := []int{20, 27, 23, 18, 24, 16, 20, 24, 18, 15, 19, 25, 24, 26,
		19, 24, 25, 21, 17, 11, 20, 21, 22, 23, 18, 20, 23, 18, 22, 23, 11}

	comfortCount(may2017) // узнаем, что было в мае 2017 г.
	comfortCount(may2018) // Узнаем, что было в мае 2018 г.
}
