package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCorrect := func (t *testing.T, sum, expected int) {
		t.Helper()

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("sum of two integers", func (t *testing.T)  {
		sum := Add(2, 2)
		expected := 4

		assertCorrect(t, sum, expected)
	})

}

// Examples to run tests
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum);
	// Output : 6
}
