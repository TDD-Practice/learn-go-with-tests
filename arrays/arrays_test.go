package arrays

import "testing"

func assertGotEqualsWant(t *testing.T, got, want int) {
	if want != got {
		t.Errorf("wanted %q, instead got %q", want, got)
	}
}

func TestArrayIterator(t *testing.T) {
	t.Run("return default for the type", func(t *testing.T) {
		got := ArrayIterator([]int{4, 5, 6})
		want := 15

		assertGotEqualsWant(t, got, want)
	})
}
