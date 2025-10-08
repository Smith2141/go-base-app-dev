package main

import "fmt"

func main() {
    minAge := 18
    maxAge := 30

    age := 17

    isValidAge := age >= minAge && age <= maxAge
    
    fmt.Println(isValidAge)
}