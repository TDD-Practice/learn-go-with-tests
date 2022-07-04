package maps

import "testing"

func assertValuesEquals(t *testing.T, got, want string) {
	if want != got {
		t.Errorf("got %q, but instead wanted %q", got, want)
	}
}

func assertErrorsEquals(t *testing.T, got, want error) {
	if want != got {
		t.Errorf("got %q, but instead wanted %q", got, want)
	}
}

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test_key": "test_value"}

	t.Run("Returns the correct value given and existing key", func(t *testing.T) {

		got, _ := dictionary.Search("test_key")
		want := "test_value"

		assertValuesEquals(t, got, want)

	})

	t.Run("Returns error for non-existing key", func(t *testing.T) {

		_, err := dictionary.Search("non_existent_key")
		want := ErrorKeyNotFound

		assertErrorsEquals(t, err, want)
	})

	t.Run("Can add a key and value pair", func(t *testing.T) {

		value, _ := dictionary.Add("new_key", "new_value")
		want := "new_value"

		assertValuesEquals(t, value, want)
	})

	t.Run(".Add() msut not ovewrite an existing value", func(t *testing.T) {
		dictionary["new_key"] = "new_value"
		_, err := dictionary.Add("new_key", "new_value")
		want := ErrorKeyAlreadyExixsts

		assertErrorsEquals(t, err, want)
	})

	t.Run("can .Update() and existing key", func(t *testing.T) {
		dictionary["new_key"] = "old_value"
		value, _ := dictionary.Update("new_key", "new_value")
		want := "new_value"

		assertValuesEquals(t, value, want)
	})

}
