package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (res [][]uint8) {
	for x:=0; x < dx; x++ {
		row := []uint8{}
		for y:=0; y < dy; y++ {
			row = append(row, uint8((x+y)/2))
		}
		res = append(res, row)
	}
	return
}

func main() {
	pic.Show(Pic)
}
