package shapes

import (
	"testing"
)

func TestPermeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// create slice with anonymous struct of tests
	areaTests := []struct{
		name string
		shape Shape
		hasArea float64
	} {
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Rectangle", shape: Rectangle{Height: 12, Width: 6}, hasArea: 72.0},
	}

	// loop over the slice and run the test cases
	for _, tt := range areaTests {
		// use name from the tt to get the name of the shape
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}