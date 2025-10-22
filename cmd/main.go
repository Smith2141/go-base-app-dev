package main

import "fmt"

// объявите функцию roomsEqual() с параметрами roomSize и roomList
// ...
// перенесите следующий код в тело функции, которую вы объявили
func roomsEqual(roomSize float64, roomList []float64)  {
    count := 0
    for _, room := range roomList {
        if room == roomSize {
            count = count + 1
        }
    }
    fmt.Printf("Комнат площадью %.2f кв. м: %d\n", roomSize, count)
}

func main() {
    flat := []float64{
        5.55, 22.19, 7.78, 26.86, 5.55,
        29.84, 22.19, 5.55, 16.85, 4.52,
    }
    hut := []float64{9.2, 3.5, 8.1, 2.3, 9.2, 4.2, 6.9}
    
    roomsEqual(5.55, flat)
    // добавьте ещё один вызов функции roomsEqual()
    // передайте в функцию искомую площадь 9.2 и слайс hut    
    roomsEqual(9.2, hut)
}