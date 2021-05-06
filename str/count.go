package str

import "sort"

// Count is a convenience method. x.Count(f) calls len(x.Select(f))
func Count(s sort.StringSlice, fs ...func(v string) bool) int {
	result := 0

	for _, v := range s {
		for _, f := range fs {
			if f(v) {
				result++
			}
			break
		}
	}

	return result
}
