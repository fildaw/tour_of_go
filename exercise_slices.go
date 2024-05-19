package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := 1; y <= dy; y++ {
		img_row := make([]uint8, dx)
		for x := 1; x <= dx; x++ {
			img_row[x-1] = uint8(x*y)
		}
		img[y-1] = img_row
	}
	return img	
}

func main() {
	pic.Show(Pic)
}
