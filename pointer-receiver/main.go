package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p Point) Add(q Point) {
	p.X += q.X
	p.Y += q.Y
}

func (p *Point) Sub(q Point) {
	p.X -= q.X
	p.Y -= q.Y
}

func main() {
	p := Point{
		X: 10,
		Y: 20,
	}
	q := Point{
		X: 1,
		Y: 2,
	}
	p.Add(q)
	fmt.Println(p) // {10 20}
	p.Sub(q)
	fmt.Println(p) // {9 18}
}
