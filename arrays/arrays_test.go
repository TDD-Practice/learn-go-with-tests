package arrays

import (
	"reflect"
	"testing"
)

func assertSliceSumGotEqualsWant(t *testing.T, got, want int) {
	if want != got {
		t.Errorf("wanted %q, instead got %q", want, got)
	}
}

func assertSliceSumAllGotEqualsWant(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %q, instead got %q", want, got)
	}
}

func TestSliceSum(t *testing.T) {
	t.Run("return default for the type", func(t *testing.T) {
		got := SliceSum([]int{4, 5, 6})
		want := 15

		assertSliceSumGotEqualsWant(t, got, want)
	})
}

func TestSliceSumAll(t *testing.T) {
	t.Run("Returns 15 for input {4,5,6}", func(t *testing.T) {
		got := SliceSumAll([]int{4, 5, 6})
		want := []int{15}

		assertSliceSumAllGotEqualsWant(t, got, want)
	})
}
