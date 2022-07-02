package structs

import "testing"

func TestRectArea(t *testing.T) {

	got := RectArea(Rectangle{2.00, 2.00})
	want := 4.00
	if want != got {
		t.Errorf("wanted %.2f but got %.2f", want, got)
	}
}

func TestRectPerim(t *testing.T) {
	got := RectPerim(Rectangle{2.00, 2.00})
	want := 8.00

	if want != got {
		t.Errorf("wanted %.2f but got %.2f", want, got)
	}
}
