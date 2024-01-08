package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat character default number of times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("repeat character specified number of times", func(t *testing.T) {
		repeated := Repeat("b", 10)
		expected := "bbbbbbbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 3)
	fmt.Println(repeated)
	// output: ccc
}
