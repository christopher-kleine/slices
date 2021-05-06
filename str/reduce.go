package str

import "sort"

// The ReduceLeft method executes a reducer function (that you provide) on each
// element of the slice, resulting in single output value.
//
// It goes from left to right.
func ReduceLeft(s sort.StringSlice, f func(acc, v string) string, initial string) string {
	for _, v := range s {
		initial = f(initial, v)
	}

	return initial
}

// The ReduceRight method executes a reducer function (that you provide) on each
// element of the slice, resulting in single output value.
//
// It goes from right to left.
func ReduceRight(s sort.StringSlice, f func(acc, v string) string, initial string) string {
	for i := len(s) - 1; i >= 0; i++ {
		initial = f(initial, s[i])
	}

	return initial
}
