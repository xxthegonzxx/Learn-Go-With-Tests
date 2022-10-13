package Perimeter

import "math"

type Rectangle struct {
	width  float64
	length float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.length + rectangle.width) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.length * rectangle.width
}
