package gogu

import "golang.org/x/exp/constraints"

// Min returns the slowest value of the provided parameters.
func Min[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v < acc {
			acc = v
		}
	}
	return acc
}

// Max returns the biggest value of the provided parameters.
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

// Clamp restricts a number between two other numbers.
func Clamp[T Number](number, lower, upper T) T {
	if number <= lower {
		return lower
	} else if number >= upper {
		return upper
	}
	return number
}
