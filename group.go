package slices

// func GroupBy[VT any, GT constraints.Ordered](slice []VT, callback func(value VT) GT) map[GT][]VT {
// 	result := make(map[GT][]VT)

// 	for _, v := range slice {
// 		k := callback(v)
// 		if _, ok := result[k]; !ok {
// 			result[k] = make([]VT, 0)
// 		}
// 		result[k] = append(result[k], v)
// 	}

// 	return result
// }
