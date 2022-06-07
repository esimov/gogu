package gogu

import (
	"golang.org/x/exp/constraints"
)

// Numbers is a custom type set of constraints extending the Float and Integer type set from the experimental constraints package.
type Numbers interface {
	constraints.Float | constraints.Integer
}

// Sum sums up all the values from the map and returns the resulted value.
func Sum[K comparable, V Numbers](m map[K]V) V {
	var acc V
	for _, v := range m {
		acc += v
	}
	return acc
}

// Invert returns a copy of the map where the keys become the values and the values the keys.
// For this to work, all of your map's values should be unique.
func Invert[K, V comparable](m map[K]V) map[V]K {
	inverted := map[V]K{}
	keys := Keys(m)

	for i := 0; i < len(keys); i++ {
		inverted[m[keys[i]]] = keys[i]
	}

	return inverted
}

// Keys retrieve all the existing keys of a map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	idx := 0
	for k, _ := range m {
		keys[idx] = k
		idx++
	}

	return keys
}

// Values retrieve all the existing values of a map.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	idx := 0
	for _, v := range m {
		values[idx] = v
		idx++
	}

	return values
}

// MapValues creates a new map with the same number of elements as the original one,
// but running each map value through an interatee function (fn).
func MapValues[K comparable, V, R any](m map[K]V, fn func(V) R) map[K]R {
	newMap := map[K]R{}

	for k, v := range m {
		newMap[k] = fn(v)
	}

	return newMap
}

// MapKeys is the opposite of MapValues. It creates a new map with the same number of elements as the original one,
// but this time the iteratee function (fn) is ran over the map keys.
func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V {
	newMap := map[R]V{}

	for k, v := range m {
		newMap[fn(k, v)] = v
	}

	return newMap
}

// Find iterates over the elements of a map and returns the first item for which the iteratee function returns true.
func Find[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V {
	var result = make(map[K]V)
	for k, v := range m {
		if fn(v) {
			result[k] = v
			break
		}
	}
	return result
}

// FindKey is like Find, but returns the first item key position for which the iteratee function returns true.
func FindKey[K comparable, V any](m map[K]V, fn func(V) bool) K {
	var result K
	for k, v := range m {
		if fn(v) {
			result = k
			break
		}
	}
	return result
}

func MapEvery[K comparable, V any](m map[K]V, fn func(V) bool) bool {
	for _, v := range m {
		if !fn(v) {
			return false
		}
	}

	return true
}

func FilterMap[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V {
	filtered := map[K]V{}

	for k, v := range m {
		if fn(v) {
			filtered[k] = v
		}
	}

	return filtered
}

func Filter2DMap[K comparable, V any](m map[K]map[K]V, fn func(V) bool) map[K]map[K]V {
	filtered := map[K]map[K]V{}

	for k, v := range m {
		for _, v1 := range v {
			if fn(v1) {
				filtered[k] = v
			}
		}
	}

	return filtered
}

// TODO consider to remove
func Filter2DMapSlice[K comparable, V any](mapSlice []map[K]map[K]V, fn func(map[K]V) bool) []map[K]map[K]V {
	filtered := []map[K]map[K]V{}

	for _, s := range mapSlice {
		for _, v := range s {
			if fn(v) {
				filtered = append(filtered, s)
			}
		}
	}

	return filtered
}
