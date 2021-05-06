package sint

import "sort"

// The All method tests whether all elements in the slice pass the test
// implemented by the provided function.
func All(s sort.IntSlice, f func(v int, i int) bool) bool {
	for i, v := range s {
		if !f(v, i) {
			return false
		}
	}

	return true
}
