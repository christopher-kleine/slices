package str

import "sort"

func IsAscSorted(s sort.StringSlice) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}

	return true
}

func IsDescSorted(s sort.StringSlice) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[i+1] {
			return false
		}
	}

	return true
}
