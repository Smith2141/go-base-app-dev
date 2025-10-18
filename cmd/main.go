package main

import "fmt"

func main() {
    currentHour := 1
    fmt.Println("На часах", currentHour, "ч.")
        
    if currentHour < 5 || currentHour >= 23 {
        fmt.Println("Доброй ночи!")
    } else if currentHour >= 5 && currentHour <= 11 {  
        fmt.Println("Доброе утро!")
    } else if currentHour >= 12 && currentHour <= 17 {  
        fmt.Println("Добрый день!")
    } else {
        fmt.Println("Добрый вечер!")
    }
}