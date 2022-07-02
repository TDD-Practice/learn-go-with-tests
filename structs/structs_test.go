package structs

import (
	"testing"
)

func TestRectArea(t *testing.T) {
	rectangle := Rectangle{2.00, 2.00}
	got := rectangle.Area()
	want := 4.00
	if want != got {
		t.Errorf("wanted %.2f but got %.2f", want, got)
	}
}

func TestRectPerim(t *testing.T) {
	rectangle := Rectangle{2.00, 2.00}
	got := rectangle.Perim()
	want := 8.00

	if want != got {
		t.Errorf("wanted %.2f but got %.2f", want, got)
	}
}

func TestPolimorphicArea(t *testing.T) {

	assertShapeArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{width: 2.00, height: 2.00}, want: 4.00},
		{shape: Circle{radius: 10.00}, want: 10.00 * 2 * 3.1416},
	}

	for _, tt := range areaTests {
		t.Run("Can calculate area of a Rectangle", func(t *testing.T) {
			assertShapeArea(t, tt.shape, tt.want)
		})
	}

	t.Run("Can calculate perimeter of a Circle", func(t *testing.T) {
		circle := Circle{10.00}
		got := circle.Perim()
		want := 10.00 * 2 * 3.1416

		if want != got {
			t.Errorf("wanted %.4f but got %.4f", want, got)
		}
	})

	/* 	t.Run("Must fail when calculating the area of a Line", func(t *testing.T) {
		line := Line{10.00, 10.00}
		assertShapeArea(t, line, 10.00*10.00*3.1416)
	}) */
}
