package main

import "fmt"

func calcRect(width, height int) (perimeter int, square int) {
    perimeter = (width+height)*2
    square = width*height
    return
}

func main() {
    perimeter, square := calcRect(5, 7)
    fmt.Println("Периметр:", perimeter, "Площадь:", square)
}