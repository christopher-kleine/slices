package str

import "sort"

// The All method tests whether all elements in the slice pass at least one of
// the teszs implemented by the provided functions.
func All(s sort.StringSlice, fs ...func(v string) bool) bool {
	for _, v := range s {
		foundIt := false
		for _, f := range fs {
			if f(v) {
				foundIt = true
				break
			}
		}
		if !foundIt {
			return false
		}
	}

	return true
}
