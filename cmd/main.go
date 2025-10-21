package main

import (
    "fmt"
    "sort"
)

func main() {
    towns := map[string]int{
        "Омск":      1104485,
        "Киров":     475464,
        "Астрахань": 465524,
        "Пятигорск": 144955,
        "Пенза2":     488299,
        "Пенза1":     488299,
        "Иркутск":   606369,
    }
    // создаём мапу с количеством жителей
    // так как кол-во жителей у городов может совпадать,
    // поэтому в качестве элементов мапы используется слайс, а не просто строка
    counts := make(map[int][]string)
    // слайс для количества жителей, который будет отсортирован
    sorted := make([]int, 0, 10)
    // формируем counts
    for name, count := range towns {
        _, ok := counts[count]
        if !ok {
            sorted = append(sorted, count)
        }
        counts[count] = append(counts[count], name)
    }
    // сортируем жителей
    sort.Ints(sorted)
    // выводим города по количеству жителей
    for _, people := range sorted {
        for _, town := range counts[people] {
            fmt.Println(town, ":", people)
        }
    }
}