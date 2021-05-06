package str_test

import (
	"testing"

	"github.com/christopher-kleine/slices/str"
)

func TestToLower(t *testing.T) {
	inp := []string{"veni", "Vedi", "viCi"}
	exp := []string{"veni", "vedi", "vici"}
	out := str.Chain(inp, str.ToLower)

	for k, v := range exp {
		if v != out[k] {
			t.Errorf("index %d missmatch: %q expected, got %q", k, v, out[k])
		}
	}
}

func TestToUpper(t *testing.T) {
	inp := []string{"veni", "Vedi", "viCi"}
	exp := []string{"VENI", "VEDI", "VICI"}
	out := str.Chain(inp, str.ToUpper)

	for k, v := range exp {
		if v != out[k] {
			t.Errorf("index %d missmatch: %q expected, got %q", k, v, out[k])
		}
	}
}
