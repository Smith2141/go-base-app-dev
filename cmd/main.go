package main

import "fmt"

// printOnTrue напечатает s в случае, если функция f вернёт true,
// в противном случае она ничего не сделает
func printOnTrue(f func(a int, b int) bool, s string) {
    if f(10, 20) {
        fmt.Println(s)
    }
}

func main() {
    // создаём анонимную функцию и присваиваем переменной f
    f := func(a, b int) bool { return a < b }
    // передаём функцию f в качестве аргумента функции printOnTrue
    printOnTrue(f, "и правда меньше")
}