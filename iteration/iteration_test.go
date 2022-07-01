package iteration

import (
	"fmt"
	"testing"
)

func assertWantEqualsGot(t *testing.T, want, got string) {
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}
}

func TestIterator(t *testing.T) {

	t.Run("When called with 'a' and 5, returns 'aaaaa' (5*'a')", func(t *testing.T) {
		got := Iterator("a", 5)
		want := "aaaaa"
		assertWantEqualsGot(t, want, got)
	})

	t.Run("When called with 'b' and 9, returns 'bbbbbbbbb' (9*'b')", func(t *testing.T) {
		got := Iterator("b", 9)
		want := "bbbbbbbbb"
		assertWantEqualsGot(t, want, got)
	})
}

func ExampleIterator() {
	fmt.Println(Iterator("a", 5))
	//Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterator("i", i)
	}
}
