package main

// добавьте в импорт необходимый пакет
import (
	"fmt"
	"strings"
)

func main() {
	message := "Алиса, в какой папке находятся мои фото?"

	path := "C:\\Documents\\Photos"

	// переведите все символы строки в переменной path в нижний регистр
	// и присвойте результат переменной lowLetterPath
	lowLetterPath := strings.ToLower(path)

	fmt.Println(message)
	fmt.Println("path =", path)

	// выведите переменную lowLetterPath и ее значение на экран
	fmt.Println("lowLetterPath =", lowLetterPath)
}
