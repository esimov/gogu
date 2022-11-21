package torx

import "golang.org/x/exp/constraints"

// Min returns the lowest value from the provided parameters.
func Min[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v < acc {
			acc = v
		}
	}
	return acc
}

// Max returns the biggest value from the provided parameters.
func Max[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v > acc {
			acc = v
		}
	}
	return acc
}

// Abs returns the absolut value of x.
func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Clamp returns a range-limited number between min and max.
func Clamp[T Number](num, min, max T) T {
	if num <= min {
		return min
	} else if num >= max {
		return max
	}
	return num
}

// InRange checks if a number is inside a range.
func InRange[T Number](num, lo, up T) bool {
	if num >= lo && num <= up {
		return true
	}
	return false
}
