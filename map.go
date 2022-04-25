package slices

func (s Slice[T]) Map(callback func(T) T) []T {
	result := make([]T, len(s))

	for index, v := range s {
		result[index] = callback(v)
	}

	return result
}

// func Map2[ST any, TT any](slice []ST, callback func(value ST, index int) TT) []TT {
// 	result := make([]TT, len(slice))

// 	for index, v := range slice {
// 		result[index] = callback(v, index)
// 	}

// 	return result
// }

// func Map3[ST any, TT any](slice []ST, callback func(value ST, index int, src []ST) TT) []TT {
// 	result := make([]TT, len(slice))

// 	for index, v := range slice {
// 		result[index] = callback(v, index, slice)
// 	}

// 	return result
// }

// func Replace[T comparable](source, target T) func(value T) T {
// 	return func(value T) T {
// 		if value == source {
// 			return target
// 		}
// 		return value
// 	}
// }
