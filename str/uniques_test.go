package str_test

import (
	"testing"

	"github.com/christopher-kleine/slices/str"
)

func TestUniques(t *testing.T) {
	inp := []string{"Roses", "are", "red", "violets", "are", "blue"}
	exp := []string{"Roses", "are", "red", "violets", "blue"}
	out := str.Chain(inp, str.Uniques)

	for k, v := range exp {
		if v != out[k] {
			t.Errorf("index %d missmatch: %q expected, got %q", k, v, out[k])
		}
	}
}
