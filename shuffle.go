package slices

func Shuffle[T any](nextInt func(int) int, values []T) []T {
	shuffled := make([]T, len(values))
	copy(shuffled, values)

	for i := len(shuffled) - 1; i > 0; i-- {
		j := nextInt(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}
