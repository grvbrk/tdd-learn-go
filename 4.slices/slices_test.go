package slices

import (
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSlices(t *testing.T) {
	t.Run("Sum a slice of integers", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		got := Sum(arr)
		want := 10
		assert(t, got, want)
	})

	t.Run("Take multiple slices as input and return the sum of each slice in a new slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Take multiple slices as input and return the sum of each slice in a new slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{11, 9}, []int{1, 3})
		want := []int{3, 9, 20, 4}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Compare 2D slices", func(t *testing.T) {
		got := Return2DSlice([]int{1, 2}, []int{0, 9}, []int{11, 9}, []int{1, 3})
		want := [][]int{{1, 2}, {0, 9}, {11, 9}, {1, 3}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

}

func assert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
