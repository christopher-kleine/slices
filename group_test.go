package slices_test

import (
	"testing"

	"github.com/christopher-kleine/slices"
)

func TestGroup(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenOdd := func(v int) float32 {
		if v%2 == 0 {
			return 0
		}
		return 1
	}
	out := slices.GroupBy(in, evenOdd)
	t.Log(out)
}
