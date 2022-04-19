package slices

func Chain[T any](slice []T, operations ...func([]T) []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	for _, op := range operations {
		result = op(result)
	}

	return result
}
