package arraysslices

import (
	"fmt"
	"reflect"
	_"slices"
	"testing"
)
func TestSum(t *testing.T) {

	t.Run("collection of five numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		// You can use reflect.DeepEqual to check for the
		// equality of the slices but it is unreliable
		// slices.Equal should be used when the slices store the same type
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
		want := []int{5, 10}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}


func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5, 56})
		SumAllTails([]int{3, 4}, []int{10, 1})
	}
}

func BenchmarkSumTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{3, 4}, []int{10, 1})
	}
}

func ExampleSum(){
	sum := Sum([]int{1, 2, 3, 4, 5, 56})
	fmt.Println(sum)
	// Output : 71
}


func ExampleSumAllTails() {
	got := SumAllTails([]int{3, 4}, []int{10, 1})
	fmt.Println(got)
	// Output : [1 1]
}