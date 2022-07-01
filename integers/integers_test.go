package integers

import (
	"fmt"
	"testing"
)

func assertWantEqualGot(t *testing.T, want, got int) {
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

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}
