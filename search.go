package slices

func (s Slice[T]) IndexOf(criteria func(T) bool) int {
	for k, v := range s {
		if criteria(v) {
			return k
		}
	}

	return -1
}

func (s Slice[T]) Contains(criteria func(T) bool) bool {
	return s.IndexOf(criteria) > -1
}

func (s Slice[T]) Find(criteria func(T) bool) *T {
	for _, v := range s {
		if criteria(v) {
			return &v
		}
	}

	return nil
}

// func Find2[T comparable](slice []T, callback func(value T, index int) bool) *T {
// 	for index, v := range slice {
// 		if callback(v, index) {
// 			return &v
// 		}
// 	}

// 	return nil
// }

// func Find3[T comparable](slice []T, callback func(value T, index int, src []T) bool) *T {
// 	for index, v := range slice {
// 		if callback(v, index, slice) {
// 			return &v
// 		}
// 	}

// 	return nil
// }
