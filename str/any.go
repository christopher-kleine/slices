package str

import "sort"

// The Any method tests whether at least one element in the slice passes the
// test implemented by the provided functions.
func Any(s sort.StringSlice, fs ...func(v string) bool) bool {
	for _, v := range s {
		for _, f := range fs {
			if f(v) {
				return true
			}
		}
	}

	return false
}
