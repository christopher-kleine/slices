package str

import "sort"

// ReverseString reverses a single string. It's UNICODE conform.
// This code was taken from Russ Cox in this post:
// https://groups.google.com/g/golang-nuts/c/oPuBaYJ17t4/m/PCmhdAyrNVkJ
func ReverseString(s sort.StringSlice) sort.StringSlice {
	for k, v := range s {
		// Get Unicode code points.
		n := 0
		rune := make([]rune, len(v))
		for _, r := range v {
			rune[n] = r
			n++
		}
		rune = rune[0:n]
		// Reverse
		for i := 0; i < n/2; i++ {
			rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
		}

		// Convert back to UTF-8.
		s[k] = string(rune)
	}

	return s
}

// ReverseOrder returns the slice in reversed order.
func ReverseOrder(s sort.StringSlice) sort.StringSlice {
	for index := 0; index < len(s)/2; index++ {
		s[index], s[len(s)-1-index] = s[len(s)-1-index], s[index]
	}

	return s
}
