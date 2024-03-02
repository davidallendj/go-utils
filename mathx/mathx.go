package mathx

import "golang.org/x/exp/constraints"

func Clamp[T constraints.Ordered](value T, min T, max T) T {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}
