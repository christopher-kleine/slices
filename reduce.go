package slices

func (s Slice[T]) Reduce(f func(currValue, reducedValue T) T, initialValue T) T {
	result := initialValue

	for _, v := range s {
		result = f(v, result)
	}

	return result
}
