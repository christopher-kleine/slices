package slices

func (s Slice[T]) Shuffle(nextInt func(int) int) []T {
	shuffled := make([]T, len(s))
	copy(shuffled, s)

	for i := len(shuffled) - 1; i > 0; i-- {
		j := nextInt(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}
