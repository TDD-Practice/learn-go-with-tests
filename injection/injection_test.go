package injection

import (
	"bytes"
	"testing"
)

func TestInjectableGreeter(t *testing.T) {
	t.Run("Can retrun a string that says 'Hello, Matheus'", func(t *testing.T) {
		buffer := bytes.Buffer{}
		InjectableGreeter(buffer, "Matheus")

		got := buffer.String()
		want := "Hello, Matheus"

		if got != want {
			t.Errorf("got %q, instead wanted %q", got, want)
		}
	})
}
