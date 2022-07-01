package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectSayHello := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Wanted %q but got %q", want, got)
		}
	}

	t.Run("When invoked with a name and lang='ENG' it says 'Hello, %name'!",
		func(t *testing.T) {
			got := sayHello("Matheus", "")
			want := engHello + ", Matheus!"
			assertCorrectSayHello(t, got, want)
		})

	t.Run("When invoked with name='' and lang='' it says 'Hello, person!'",
		func(t *testing.T) {
			got := sayHello("", "")
			want := engHello + ", person!"
			assertCorrectSayHello(t, got, want)
		})

	t.Run("When invoked with a name and lang='ESP' it says 'Hola, %name'!", func(t *testing.T) {
		got := sayHello("Matheus", "ESP")

		want := espHello + ", Matheus!"
		assertCorrectSayHello(t, got, want)
	})

	t.Run("When invoked with a name and lang='FRA' it says 'Bonjour, %name'!", func(t *testing.T) {
		got := sayHello("Matheus", "ESP")

		want := espHello + ", Matheus!"
		assertCorrectSayHello(t, got, want)
	})

}
