package main 

import "fmt"

func main() {
    list := []int{1, 5, 4, 8, 7, 6, 2, 3}
	newlist := list
	fmt.Println(newlist)

    newlist[0] = 77
    newlist[len(newlist)-1] = 99

    fmt.Println(newlist)
    fmt.Println(list)
}