package slices_test

import (
	"testing"

	"github.com/christopher-kleine/slices"
)

func TestReduceSimple(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := func(cv, tv int) int {
		return cv + tv
	}
	out := slices.Slice[int](in).Reduce(sum, 0)
	exp := 55
	if out != exp {
		t.Errorf("%d != %d", out, exp)
	}
}

func TestReduceComplex(t *testing.T) {
	type Record struct {
		Int   int
		Float float32
	}

	initialValue := Record{
		Int:   0,
		Float: 0,
	}
	in := []Record{
		{Int: 1, Float: 2},
		{Int: 2, Float: 4},
		{Int: 3, Float: 6},
		{Int: 4, Float: 8},
		{Int: 5, Float: 10},
		{Int: 6, Float: 12},
		{Int: 7, Float: 14},
		{Int: 8, Float: 16},
		{Int: 9, Float: 18},
		{Int: 10, Float: 20},
	}
	sum := func(cv, tv Record) Record {
		tv.Int = tv.Int + cv.Int
		return tv
	}
	out := slices.Slice[Record](in).Reduce(sum, initialValue)
	t.Log(out)
	t.Log(initialValue)
	t.Log(in)
}
