package str

import (
	"sort"
	"strings"
)

// HasPrefix returns only the entries with the given Prefixes
func HasPrefix(prefixes ...string) func(string) bool {
	return func(v string) bool {
		for _, prefix := range prefixes {
			if strings.HasPrefix(v, prefix) {
				return true
			}
		}

		return false
	}
}

// AddPrefix adds a prefix to every entry in the StringSlice
func AddPrefix(prefix string) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		for index := range src {
			src[index] = prefix + src[index]
		}

		return src
	}
}

// TrimPrefix removes a prefix from every entry in the StringSlice
func TrimPrefix(prefix string) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		for index := range src {
			src[index] = strings.TrimPrefix(src[index], prefix)
		}

		return src
	}
}
