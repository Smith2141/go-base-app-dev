package main

import "fmt"

func main() {
    // создаём анонимную функцию и присваиваем переменной f
    f := func(a, b int) bool { return a < b }
    // вызываем функцию и сохраняем результат в переменную res
    res := f(10, 20)                         

    fmt.Println(res)
}