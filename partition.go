package slices

func Partition[T comparable](slice []T, isLeft func(value T) bool) (left, right []T) {
	for _, v := range slice {
		if isLeft(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return left, right
}

func Partition2[T comparable](slice []T, isLeft func(value T, index int) bool) (left, right []T) {
	for index, v := range slice {
		if isLeft(v, index) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return left, right
}

func Partition3[T comparable](slice []T, isLeft func(value T, index int, src []T) bool) (left, right []T) {
	for index, v := range slice {
		if isLeft(v, index, slice) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return left, right
}
