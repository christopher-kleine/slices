package str

import "sort"

// Partition returns two new slices based on the test-function. If the
// test-function returns true, the element will be included in the first slice,
// in the second slice otherwise.
func Partition(s sort.StringSlice, f func(v string) bool) (sort.StringSlice, sort.StringSlice) {
	left := make(sort.StringSlice, len(s))
	right := make(sort.StringSlice, len(s))
	leftCount := 0
	rightCount := 0

	for _, v := range s {
		if f(v) {
			left[leftCount] = v
			leftCount++
		} else {
			right[leftCount] = v
			rightCount++
		}
	}

	return left[:leftCount], right[:rightCount]
}
