package main

import "fmt"

// объявите функцию здесь
func printFriendsCount(friendsCount int) {
	// код, написанный ниже, переместите внутрь объявленной вами функции
	if friendsCount == 0 {
		fmt.Println("У тебя нет друзей")
	} else if friendsCount == 1 {
		fmt.Println("У тебя", friendsCount, "друг")
	} else if friendsCount >= 2 && friendsCount <= 4 {
		fmt.Println("У тебя", friendsCount, "друга")
	} else if friendsCount >= 5 && friendsCount < 20 {
		fmt.Println("У тебя", friendsCount, "друзей")
	} else {
		fmt.Println("Ого, сколько у тебя друзей! Целых ", friendsCount)
	}
}

// вызовите объявленную вами функцию printFriendsCount()
// с нужными аргументами внутри функции main
func main() {
	for i := 0; i <= 20; i++ {
		printFriendsCount(i)
	}
}
