package str

import "sort"

// Chain takes a source slice and slice-operations. It returns a new slice with
// all operations applied to it.
func Chain(src sort.StringSlice, operations ...func(sort.StringSlice) sort.StringSlice) sort.StringSlice {
	result := make([]string, len(src))
	copy(result, src)

	for _, operation := range operations {
		result = operation(result)
	}

	return result
}
