package str

import "sort"

func Duplicates(s sort.StringSlice) sort.StringSlice {
	lookUp := make(map[string]bool)
	result := make(sort.StringSlice, 0)

	for _, v := range s {
		if _, ok := lookUp[v]; !ok {
			lookUp[v] = true
		}
	}

	for _, v := range s {
		if _, ok := lookUp[v]; !ok {
			lookUp[v] = false
			result = append(result, v)
		}
	}

	return result
}
