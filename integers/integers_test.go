package integers_test

import (
	"testing"
)

func assertWantEqualGot(t testing.T, got, want int) {
	if want != got {
		t.Errorf("want %d instead got %d", want, got)
	}
}

func TestAdder(t *testing.T) {
	t.Run("first failing test", func(t *testing.T) {
		got := Adder(1, 2)
		want := 3
		assertWantEqualGot(t, want, got)

	})
}
