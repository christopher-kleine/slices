package str

import "sort"

func GroupBy(s sort.StringSlice, f func(v string) string) map[string]sort.StringSlice {
	result := make(map[string]sort.StringSlice)

	for _, v := range s {
		result[f(v)] = append(result[f(v)], v)
	}

	return result
}
