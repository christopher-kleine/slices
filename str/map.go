package str

import "sort"

// Map takes converts a StringSlice based on the map-function provided as
// argument
func Map(f func(string) string) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		for index, value := range src {
			src[index] = f(value)
		}

		return src
	}
}
