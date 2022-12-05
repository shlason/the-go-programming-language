package geometry

import "math"

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (c Point) Distance(q Point) float64 {
	return math.Hypot(q.X-c.X, q.Y-c.Y)
}
