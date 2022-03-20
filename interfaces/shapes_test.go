package interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := Perimeter(r)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f hasArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// The %#v format string will print out our struct with the values in its field
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
