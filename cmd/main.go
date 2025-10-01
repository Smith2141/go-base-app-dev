package main

import "fmt"

const (
    a = 1
    b
    c = 5
    d
)

func main() {
    fmt.Println(a, b, c, d) // 1 1 5 5
}