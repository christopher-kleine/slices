package slices

func (s Slice[T]) Select(criteria ...func(value T) bool) []T {
	result := []T{}

	for _, v := range s {
		for _, callback := range criteria {
			if callback(v) {
				result = append(result, v)
				break
			}
		}
	}

	return result
}

// func Select2[T constraints.Ordered](slice []T, comp func(value T, index int) bool) []T {
// 	result := []T{}

// 	for index, v := range slice {
// 		if comp(v, index) {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

// func Select3[T constraints.Ordered](slice []T, comp func(value T, index int, src []T) bool) []T {
// 	result := []T{}

// 	for index, v := range slice {
// 		if comp(v, index, slice) {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

func (s Slice[T]) Reject(criteria ...func(value T) bool) []T {
	result := []T{}

	for _, v := range s {
		keepIt := true
		for _, callback := range criteria {
			if callback(v) {
				keepIt = false
				break
			}
		}

		if keepIt {
			result = append(result, v)
		}
	}

	return result
}

// func Reject2[T constraints.Ordered](slice []T, comp func(value T, index int) bool) []T {
// 	result := []T{}

// 	for index, v := range slice {
// 		if !comp(v, index) {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

// func Reject3[T constraints.Ordered](slice []T, comp func(value T, index int, source []T) bool) []T {
// 	result := []T{}

// 	for index, v := range slice {
// 		if !comp(v, index, slice) {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

// func SelectC[T constraints.Ordered](criteria ...func(value T) bool) func(slice []T) []T {
// 	result := []T{}

// 	return func(slice []T) []T {
// 		for _, v := range slice {
// 			for _, callback := range criteria {
// 				if callback(v) {
// 					result = append(result, v)
// 					break
// 				}
// 			}
// 		}

// 		return result
// 	}
// }

// func Select2C[T constraints.Ordered](comp func(value T, index int) bool) func(slice []T) []T {
// 	result := []T{}

// 	return func(slice []T) []T {
// 		for index, v := range slice {
// 			if comp(v, index) {
// 				result = append(result, v)
// 			}
// 		}

// 		return result
// 	}
// }

// func Select3C[T constraints.Ordered](comp func(value T, index int, src []T) bool) func(slice []T) []T {
// 	result := []T{}

// 	return func(slice []T) []T {
// 		for index, v := range slice {
// 			if comp(v, index, slice) {
// 				result = append(result, v)
// 			}
// 		}

// 		return result
// 	}
// }

// func RejectC[T constraints.Ordered](slice []T, comp func(value T) bool) []T {
// 	result := []T{}

// 	for _, v := range slice {
// 		if !comp(v) {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

// func Reject2C[T constraints.Ordered](slice []T, comp func(value T, index int) bool) []T {
// 	result := []T{}

// 	for index, v := range slice {
// 		if !comp(v, index) {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

// func Reject3C[T constraints.Ordered](slice []T, comp func(value T, index int, source []T) bool) []T {
// 	result := []T{}

// 	for index, v := range slice {
// 		if !comp(v, index, slice) {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }
