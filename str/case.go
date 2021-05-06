package str

import (
	"sort"
	"strings"
)

// ToUpper uses "strings.ToUpper" on every entry in the StringSlice
func ToUpper(src sort.StringSlice) sort.StringSlice {
	for index := range src {
		src[index] = strings.ToUpper(src[index])
	}

	return src
}

// ToUpper uses "strings.ToLower" on every entry in the StringSlice
func ToLower(src sort.StringSlice) sort.StringSlice {
	for index := range src {
		src[index] = strings.ToLower(src[index])
	}

	return src
}
