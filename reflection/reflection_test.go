package reflection

import "testing"

func TestWalk(t *testing.T) {

	expected := "Matheus"

	var received []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(s string) {
		received = append(received, s)
	})

	if len(received) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(received), 1)
	}
}
