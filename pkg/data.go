package data

import "math"

var x int = 1

func GetX() int {
	return x
}

type Point struct {
	x, y float64
}

func (p *Point) InitMe(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) OffsetX(x float64) {
	p.x = p.x + x
}

func (p *Point) DistToOrig() float64 {
	t := (p.x * p.x) + (p.y * p.y)
	return math.Sqrt(t)
}

func (p *Point) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}
