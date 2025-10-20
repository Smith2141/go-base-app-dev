package main

import "fmt"

func main() {
	words := []string{"Терры", "важны", "в", "Роботы", "стали",
		"период", "эмиграции", "c"}

	sentence := make([]string, 0, 8)
	sentence = append(sentence, words[3:5]...)
	// добавьте в правильном порядке остальные три фрагмента
	sentence = append(sentence, words[1:3]...)
	sentence = append(sentence, words[5:]...)
	sentence = append(sentence, words[0])

	fmt.Println(sentence)
}
