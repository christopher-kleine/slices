package slices

import "strings"

func TrimSuffix(suffix string) func(s string) string {
	return func(s string) string {
		return strings.TrimSuffix(s, suffix)
	}
}

func TrimPrefix(prefix string) func(s string) string {
	return func(s string) string {
		return strings.TrimPrefix(s, prefix)
	}
}

func AddSuffix(suffix string) func(s string) string {
	return func(s string) string {
		return s + suffix
	}
}

func AddPrefix(prefix string) func(s string) string {
	return func(s string) string {
		return prefix + s
	}
}
