package main

// Notice: don't forget to `go get golang.org/x/tour/pic`

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([]uint8, dy)

	for i, _ := range y {
		
	}

	return y
}

func main() {
	pic.Show(Pic)
}
