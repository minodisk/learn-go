package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func NewVertex(x, y float64) *Vertex {
	return &Vertex{x, y}
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	v = NewVertex(1, math.Sqrt(3))
	fmt.Println(v.Abs())
}
