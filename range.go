package gogu

import (
	"errors"
)

// Range creates a slice of numbers (integers) progressing from start (if omitted defaults to 0) until the end.
// This method can accept 1, 2 or 3 parameters. Depending on the number of provided parameters, `start`, `step` and `end` are having the following meanings:
// [start=0]: The start of the range.
// [step=1]: The value to increment or decrement by.
// end: The end of the range.

// In case you'd like negative values, use a negative step.
// TODO make a thorough test.
func Range[T ~int](params ...T) ([]T, error) {
	var result []T

	if len(params) > 3 {
		return nil, errors.New("the method require maximum 3 paramenters")
	}

	var start, step, end T

	switch len(params) {
	case 1:
		step = 1
		end = params[len(params)-1]
	case 2:
		start = params[0]
		step = 1
		end = params[len(params)-1]
	case 3:
		start = params[0]
		step = params[1]
		end = params[len(params)-1]

		if step == 0 {
			return nil, errors.New("step value should not be zero")
		}
		if step < 0 && end > start {
			return nil, errors.New("the end value should be less than the start value in case you are using a negative increment")
		}
	default:
		return nil, errors.New("the method require at least one paramenter, which should be the range dimension in this case")
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
func RangeRight[T ~int](params ...T) ([]T, error) {
	ran, err := Range(params...)
	if err != nil {
		return nil, err
	}
	return Reverse(ran), nil
}
