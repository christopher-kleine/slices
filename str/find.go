package str

import "sort"

func FindValue(s sort.StringSlice, fs ...func(v string) bool) (string, bool) {
	for i, v := range s {
		for _, f := range fs {
			if f(v) {
				return s[i], true
			}
		}
	}

	return "", false
}

func FindIndex(s sort.StringSlice, fs ...func(v string) bool) (int, bool) {
	for i, v := range s {
		for _, f := range fs {
			if f(v) {
				return i, true
			}
		}
	}

	return 0, false
}
