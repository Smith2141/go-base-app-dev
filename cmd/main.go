package main

import "fmt"

func main() {
    i := 5
    switch {
    case i == 0:
        fmt.Println("o")
        fallthrough
    case i < 5, i > 4:
        fmt.Println("oo")
    case i < 10:
        fmt.Println("ooo")
    case i == 5:
        fmt.Println("oooo")
    default: 
        fmt.Println("ooooo")
    }
}