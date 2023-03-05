package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	array := make([][]uint8, dy)
	for y := range array {
		array[y] = make([]uint8, dx)
		for x := range array[y] {
			array[y][x] = uint8(x ^ y)
		}
	}
	return array
}

func main() {
	pic.Show(Pic)
}
