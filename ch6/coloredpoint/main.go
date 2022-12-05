package main

import (
	"fmt"
	"image/color"
)

type Count int

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	*Count
	color.RGBA
}

func (c *Count) Write(n int) {
	*c = Count(n)
}

func main() {
	c := Count(1)
	cp := ColoredPoint{Point{1, 2}, &c, color.RGBA{1, 1, 1, 1}}
	fmt.Println(*cp.Count)
	c.Write(100)
}
