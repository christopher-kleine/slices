package str

import "sort"

func Min(s sort.StringSlice) string {
	if len(s) == 0 {
		return ""
	}
	min := s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func Max(s sort.StringSlice) string {
	if len(s) == 0 {
		return ""
	}
	max := s[0]

	for _, v := range s {
		if v > max {
			max = v
		}
	}

	return max
}
