package slices_test

import (
	"testing"

	"github.com/christopher-kleine/slices"
)

func TestChain(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	out := slices.Chain(in,
		slices.SelectC(slices.DividableBy(3), slices.DividableBy(2)),
		slices.SelectC(slices.DividableBy(5)),
	)
	t.Log(out)
}
