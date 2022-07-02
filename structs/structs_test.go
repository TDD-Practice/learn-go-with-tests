package structs

import "testing"

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
