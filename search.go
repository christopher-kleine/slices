package slices

func IndexOf[T comparable](slice []T, value T) int {
	for k, v := range slice {
		if v == value {
			return k
		}
	}

	return -1
}

func Contains[T comparable](slice []T, value T) bool {
	return IndexOf(slice, value) > -1
}

func Find[T comparable](slice []T, callback func(value T) bool) *T {
	for _, v := range slice {
		if callback(v) {
			return &v
		}
	}

	return nil
}

func Find2[T comparable](slice []T, callback func(value T, index int) bool) *T {
	for index, v := range slice {
		if callback(v, index) {
			return &v
		}
	}

	return nil
}

func Find3[T comparable](slice []T, callback func(value T, index int, src []T) bool) *T {
	for index, v := range slice {
		if callback(v, index, slice) {
			return &v
		}
	}

	return nil
}
