package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base float64
}

// Area of a Rectangle
func (rect Rectangle) Area() float64 {
	return (rect.Height * rect.Width)
}

// Area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area of a triange
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func Perimeter(rect Rectangle) float64 {
	return  2 * (rect.Height + rect.Width)
}
