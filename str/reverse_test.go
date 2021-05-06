package str_test

import (
	"testing"

	"github.com/christopher-kleine/slices/str"
)

func TestReverseString(t *testing.T) {
	inp := []string{"Hello", "World", "N(ò,ó)ß", "Gôtherfäüst"}
	exp := []string{"olleH", "dlroW", "ß)ó,ò(N", "tsüäfrehtôG"}
	out := str.Chain(inp, str.ReverseString)

	for k, v := range exp {
		if v != out[k] {
			t.Errorf("index %d missmatch: %q expected, got %q", k, v, out[k])
		}
	}
}

func TestReverseOrderEven(t *testing.T) {
	inp := []string{"Roses", "are", "red", "violets", "are", "blue"}
	exp := []string{"blue", "are", "violets", "red", "are", "Roses"}
	out := str.Chain(inp, str.ReverseOrder)

	for k, v := range exp {
		if v != out[k] {
			t.Errorf("index %d missmatch: %q expected, got %q", k, v, out[k])
		}
	}
}

func TestReverseOrderOdd(t *testing.T) {
	inp := []string{"Roses", "are", "red", "violets", "are", "blue", "!"}
	exp := []string{"!", "blue", "are", "violets", "red", "are", "Roses"}
	out := str.Chain(inp, str.ReverseOrder)

	for k, v := range exp {
		if v != out[k] {
			t.Errorf("index %d missmatch: %q expected, got %q", k, v, out[k])
		}
	}
}
