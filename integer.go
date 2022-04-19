package slices

import "golang.org/x/exp/constraints"

func Add[T constraints.Integer](v T) func(value T) T {
	return func(value T) T {
		return v + value
	}
}

func Sub[T constraints.Integer](v T) func(value T) T {
	return func(value T) T {
		return v - value
	}
}

func Mul[T constraints.Integer](v T) func(value T) T {
	return func(value T) T {
		return v * value
	}
}

func Div[T constraints.Integer](v T) func(value T) T {
	return func(value T) T {
		return value / v
	}
}

func Mod[T constraints.Integer](v T) func(value T) T {
	return func(value T) T {
		return value % v
	}
}

func DividableBy[T constraints.Integer](div T) func(value T) bool {
	return func(value T) bool {
		return value%div == 0
	}
}
