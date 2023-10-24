package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)
	fmt.Println(p1.Length(p2))

}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Length(p2 Point) float64 {
	dx := p.x - p2.GetX()
	dy := p.y - p2.GetY()
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}
