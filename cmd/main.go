package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string = "21"
	var y string = "25.4"

	girlfriendAge, err := strconv.Atoi(x)

	if err != nil {
		panic(err)
	}

	myAge, err := strconv.ParseFloat(y, 64)

	if err != nil {
		panic(err)
	}

    var sumOfAges int = girlfriendAge + int(myAge)

	fmt.Println(sumOfAges)
}
