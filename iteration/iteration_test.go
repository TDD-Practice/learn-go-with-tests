package iteration

import "testing"

func assertWantEqualsGot(t *testing.T, want, got string) {
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}
}

func TestIterator(t *testing.T) {
	t.Run("Returns the zero value for the type when called with empty parameter", func(t *testing.T) {
		got := Iterator("")
		want := ""
		assertWantEqualsGot(t, want, got)
	})

	t.Run("When called with 'a' returns 'aaaaa' (5*'a')", func(t *testing.T) {
		got := Iterator("a")
		want := "aaaa"
		assertWantEqualsGot(t, want, got)
	})
}
