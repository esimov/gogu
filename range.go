package gogu

import (
	"errors"
)

// Range creates a slice of numbers (integers) progressing from start up to, but not including end.
// This method can accept 1, 2 or 3 arguments.
// Depending on the number of provided parameters, `start`, `step` and `end` has the following meaning:
// [start=0]: The start of the range. If ommited it defaults to 0.
// [step=1]: The value to increment or decrement by.
// end: The end of the range.

// In case you'd like negative values, use a negative step.
func Range[T Number](args ...T) ([]T, error) {
	var result []T

	if len(args) > 3 {
		return nil, errors.New("the method require maximum 3 paramenters")
	}

	var start, step, end T

	switch len(args) {
	case 1:
		step = 1
		end = args[len(args)-1]
	case 2:
		start = args[0]
		step = 1
		end = args[len(args)-1]
	case 3:
		start = args[0]
		step = args[1]
		end = args[len(args)-1]

		if start > end && end > 0 {
			return nil, errors.New("the end value should be greater than start value")
		}
		if step == 0 {
			return nil, errors.New("step value should not be zero")
		}
		if step < 0 && end > start {
			return nil, errors.New("the end value should be less than the start value in case you are using a negative increment")
		}
	}

	if end > 0 {
		for i := start; i < end; i += step {
			result = append(result, i)
		}
	} else {
		for i := start; end < i; i -= Abs(step) {
			result = append(result, i)
		}
	}

	return result, nil
}

// RangeRight is like Range, only that it populates the slice in descending order.
func RangeRight[T Number](params ...T) ([]T, error) {
	ran, err := Range(params...)
	if err != nil {
		return nil, err
	}
	return Reverse(ran), nil
}
