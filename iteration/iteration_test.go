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

	t.Run("When called with 'a' returns 'aaaaa' (5*'a')", func(t *testing.T) {
		got := Iterator("a")
		want := "aaaaa"
		assertWantEqualsGot(t, want, got)
	})
}

func ExampleIterator() {
	fmt.Println(Iterator("a"))
	//Output: aaaaa
}
