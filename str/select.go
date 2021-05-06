package str

import "sort"

// Select removes all entries that do not fullfill at least one of the
// provided validation functions.
func Select(fs ...func(string) bool) func(sort.StringSlice) sort.StringSlice {
	return func(src sort.StringSlice) sort.StringSlice {
		cache := make(map[string]bool)
		for _, v := range src {
			if _, ok := cache[v]; !ok {
				cache[v] = false
				for _, fn := range fs {
					if res := fn(v); res {
						cache[v] = true
						break
					}
				}
			}
		}

		for index := 0; index < len(src); index++ {
			if !cache[src[index]] {
				src = append(src[:index], src[index+1:]...)
			}
		}

		return src
	}
}
