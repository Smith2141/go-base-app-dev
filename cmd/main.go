package main

import "fmt"

func main() {
    // рекомендованные фильмы
    recommended := [...]string{"Хатико", "23", "Достучаться до небес",
        "Хакеры", "Трон", "1408"}
    // коллекция
    collection := [...]string{"Трон", "Военные игры", "Тихушники",
        "Джонни Мнемоник", "Хакеры", "Нирвана", "23", "Враг государства",
        "Взлом", "Пароль рыба-меч", "Сеть", "Кто я"}

	for recomendation := range recommended {
		for movie := range collection {
			if recommended[recomendation] == collection[movie] {
				fmt.Println(collection[movie])
			}
		}
	}
}
