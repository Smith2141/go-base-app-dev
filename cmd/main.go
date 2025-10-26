package main

import "fmt"

// вместо многоточия добавьте тип возвращаемого значения int
func comfortCount(temperatures []int) int {
	count := 0
	for _, temp := range temperatures {
		if temp >= 22 && temp <= 26 {
			count += 1
		}
	}
	// добавьте return со значением
	return count
}

func main() {
	may2017 := []int{24, 26, 15, 10, 15, 19, 10, 1, 4, 7, 7, 7, 12, 14,
		17, 8, 9, 19, 21, 22, 11, 15, 19, 23, 15, 21, 16, 13, 25, 17, 19}

	// код ниже не изменяйте:
	// вызовем функцию comfortCount(), передадим в неё слайс may2017,
	// результат работы сохраним в переменную niceDays
	niceDays := comfortCount(may2017)

	// напечатаем значение, сохранённое в niceDays
	fmt.Println("Количество тёплых дней в этом месяце:", niceDays)
}
