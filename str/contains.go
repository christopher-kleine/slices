package str

import (
	"sort"
	"strings"
)

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

// The Contains function returns a validate-function. This function checks if
// the provided string contains at least one of the provided values.
func Contains(values ...string) func(v string) bool {
	return func(v string) bool {
		for _, value := range values {
			if strings.Contains(v, value) {
				return true
			}
		}

		return false
	}
}
