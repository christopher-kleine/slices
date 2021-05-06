package str

import "sort"

// Reject removes all entrys in the StringSlice that pass the validation
// function
func Reject(fs ...func(string) bool) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		result := make(sort.StringSlice, 0)
		for _, v := range src {
			for _, f := range fs {
				if !f(v) {
					result = append(result, v)
					break
				}
			}
		}

		return result
	}
}
