package str

import (
	"sort"
	"strings"
)

// HasSuffix checks if a string has a number of suffixes
func HasSuffix(suffixes ...string) func(string) bool {
	return func(v string) bool {
		for _, suffix := range suffixes {
			if strings.HasSuffix(v, suffix) {
				return true
			}
		}

		return false
	}
}

// AddSuffix adds a suffix to every entry in the StringSlice
func AddSuffix(suffix string) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		for index := range src {
			src[index] = src[index] + suffix
		}

		return src
	}
}

// TrimSuffix removes a suffix from every entry in the StringSlice
func TrimSuffix(prefix string) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		for index := range src {
			src[index] = strings.TrimSuffix(src[index], prefix)
		}

		return src
	}
}
