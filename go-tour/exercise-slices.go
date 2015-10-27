/*
 * exercise-slices.go
 */

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	canvas := make([][]uint8, dy)
	for i := range canvas {
		canvas[i] = make([]uint8, dx)
	}

	for i := range canvas {
		for j := range canvas[i] {
			canvas[i][j] = func(x, y int) uint8 { return uint8(x + y) }(j, i)
		}
	}

	return canvas
}

func main() {
	pic.Show(Pic)
}
