package main

import (
	"fmt"
)

type Abc float64

type TestT struct {
	X, Y int
}

type Point struct {
	TestT
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Spokes struct {
	X, Y, Num int
}

type Wheel struct {
	Circle
	Spokes
	Abc
}

func main() {
	w := Wheel{
		Circle: Circle{
			Point: Point{
				X: 5,
				Y: 10,
			},
			Radius: 20,
		},
		Spokes: Spokes{
			X:   9,
			Y:   18,
			Num: 36,
		},
	}
	w1 := Wheel{Circle{Point{TestT{1, 2}, 3, 4}, 20}, Spokes{9, 18, 36}, 1.1}

	fmt.Printf("%#v\t%#v\n", w, w1.Abc)
}
