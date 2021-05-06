package str

import "sort"

// Remove delets all entries in the StringSlice that match the given string
// values. This is shorter and faster than using the "Reject" function.
func Remove(values ...string) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		result := make(sort.StringSlice, 0)
		valueMap := make(map[string]bool)

		for _, value := range values {
			valueMap[value] = true
		}

		for _, value := range src {
			if _, ok := valueMap[value]; !ok {
				result = append(result, value)
			}
		}

		return result
	}
}
