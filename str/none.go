package str

import "sort"

// The None method tests whether no element in the slice passes the
// test implemented by the provided function.
func None(s sort.StringSlice, fs ...func(v string) bool) bool {
	for _, v := range s {
		for _, f := range fs {
			if f(v) {
				return false
			}
		}
	}

	return true
}
