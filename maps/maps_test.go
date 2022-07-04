package maps

import "testing"

func assertValuesEquals(t *testing.T, got, want string) {
	if want != got {
		t.Errorf("got %q, but instead wanted %q", got, want)
	}
}

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test_key": "test_value"}

	t.Run("Can read a value from a map given and existing key", func(t *testing.T) {

		got, err := dictionary.Search("test_key")
		want := "test_value"

		assertValuesEquals(t, got, want)

	})

	t.Run("Error if looking for non-existing key", func(t *testing.T) {

		got, err := dictionary.Search("no_existent_key")
		want := "test_value"

		assertValuesEquals(t, got, want)
	})

}
