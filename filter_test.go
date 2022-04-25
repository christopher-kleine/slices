package slices_test

import (
	"testing"

	"github.com/christopher-kleine/slices"
)

// func TestSelect(t *testing.T) {
// 	intSlice := []int{0, 2, 3, 4, 5, 6, 7, 8, 9}

// }

func TestReject(t *testing.T) {
	in := []string{"Atom", "Brackets", "Code", "Disco", "electron", "final"}
	out := slices.Slice[string](in).Reject(func(value string) bool {
		return value[0] == 'D'
	}, func(value string) bool {
		return value[0] == 'e'
	})
	exp := []string{"Atom", "Brackets", "Code", "final"}

	for k, v := range out {
		if v != exp[k] {
			t.Errorf("%s != %s (Index %d)", v, exp[k], k)
		}
	}
}
