package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("When we want to sum a collection", func(t *testing.T) {
		t.Run("And is a collection of 'N' number", func(t *testing.T) {
			numbers := []int{1, 2}

			got := Sum(numbers)
			want := 3

			if got != want {
				t.Errorf("got %d want %d given, %d", got, want, numbers)
			}
		})
	})
}

func checkSum(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSumAllTrails(t *testing.T) {

	t.Run("When we sum tails of collections", func(t *testing.T) {
		t.Run("And there's an empty collection", func(t *testing.T) {
			got := SumAllTrails([]int{}, []int{3, 4, 5})
			want := []int{0, 9}

			checkSum(t, got, want)
		})

		t.Run("And all collections have non-empty value", func(t *testing.T) {
			got := SumAllTrails([]int{1, 2}, []int{0, 9})
			want := []int{2, 9}

			checkSum(t, got, want)
		})
	})
}
