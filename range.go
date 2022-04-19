package slices

import (
	"golang.org/x/exp/constraints"
)

func Below[T constraints.Ordered](bottom T, includeEqual bool) func(value T) bool {
	return func(value T) bool {
		return value < bottom || (includeEqual && value == bottom)
	}
}

func Above[T constraints.Ordered](bottom T, includeEqual bool) func(value T) bool {
	return func(value T) bool {
		return value > bottom || (includeEqual && value == bottom)
	}
}

func Between[T constraints.Ordered](bottom, roof T, includeEqual bool) func(value T) bool {
	above := Above(bottom, includeEqual)
	below := Below(roof, includeEqual)

	return func(value T) bool {
		return above(value) && below(value)
	}
}
