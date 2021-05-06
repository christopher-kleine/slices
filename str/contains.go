package str

import "sort"

// The ContainsValue method determines whether the slice includes certain values
// among its entries, returning true or false as appropriate.
func ContainsValue(s sort.StringSlice, values ...string) bool {
	lookUp := make(map[string]bool)
	for _, v := range values {
		lookUp[v] = false
	}

	for _, v := range s {
		if _, ok := lookUp[v]; ok {
			return true
		}
	}

	return false
}
