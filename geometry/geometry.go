package geometry

import "math"

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

func (p Path) Distance() float64 {
	d := 0.0
	for i := range p {
		if i > 0 {
			d += p[i].Distance(p[i-1])
		}
	}
	return d
}
