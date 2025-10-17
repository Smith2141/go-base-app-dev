package main

import "fmt"

const needVolume int32 = 5_000_000

func main() {
	var minSize int32 = 1_00
	var maxSize int32 = 3_00

main:
	for i := minSize; i < maxSize; i++ {
		var square int32 = i * i
		for j := minSize; j < maxSize; j++ {
			var volume int32 = square * j
			if volume == needVolume {
				fmt.Println(i, i, j, volume)
				break main
			}
		}
	}
}
