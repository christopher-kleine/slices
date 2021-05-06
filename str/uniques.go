package str

import "sort"

// Uniques removes all doubles from the StringSlice. The first entry however
// is kept.
func Uniques(s sort.StringSlice) sort.StringSlice {
	lookUp := make(map[string]bool)

	i := 0
	for i < len(s) {
		if _, ok := lookUp[s[i]]; !ok {
			lookUp[s[i]] = true
			i++
		} else {
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}
