package gogu

import (
	"golang.org/x/exp/constraints"
)

// Number is a custom type set of constraints extending the Float and Integer type set from the experimental constraints package.
type Number interface {
	constraints.Float | constraints.Integer
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
// but running each map value through a callback function (fn).
func MapValues[K comparable, V, R any](m map[K]V, fn func(V) R) map[K]R {
	newMap := map[K]R{}

	for k, v := range m {
		newMap[k] = fn(v)
	}

	return newMap
}

// MapKeys is the opposite of MapValues. It creates a new map with the same number of elements
// as the original one, but this time the callback function (fn) is invoked over the map keys.
func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V {
	newMap := map[R]V{}

	for k, v := range m {
		newMap[fn(k, v)] = v
	}

	return newMap
}

// Find iterates over the elements of a map and returns the first item for which the callback function returns true.
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

// FindKey is like Find, but returns the first item key position for which the callback function returns true.
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

// FindByKey is like Find, but returns the first item key position for which the callback function returns true.
func FindByKey[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V {
	var result = make(map[K]V)
	for k, v := range m {
		if fn(k) {
			result[k] = v
		}
	}
	return result
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

// MapEvery returns true if all of the elements of a map satisfies the criteria of the callback function.
func MapEvery[K comparable, V any](m map[K]V, fn func(V) bool) bool {
	for _, v := range m {
		if !fn(v) {
			return false
		}
	}

	return true
}

// MapSome returns true if some of the elements of a map satisfies the criteria of the callback function.
func MapSome[K comparable, V any](m map[K]V, fn func(V) bool) bool {
	for _, v := range m {
		if fn(v) {
			return true
		}
	}

	return false
}

// MapContains returns true if the value is present in the list otherwise false.
func MapContains[K, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// MapCollection is like the Map method applied on slices, but this time applied on maps.
// It runs each element over an iteratee function and saves the resulted values into a new map.
func MapCollection[K comparable, V any](m map[K]V, fn func(V) V) []V {
	result := make([]V, len(m))

	idx := 0
	for _, v := range m {
		result[idx] = fn(v)
		idx++
	}
	return result
}

// Pluck extracts all the values of a map by the key definition.
func Pluck[K comparable, V any](mapSlice []map[K]V, key K) []V {
	var result = []V{}

	for _, m := range mapSlice {
		mapped := FindByKey(m, func(k K) bool {
			return k == key
		})
		if _, ok := mapped[key]; ok {
			result = append(result, mapped[key])
		}
	}

	return result
}

// Pick extracts the elements from the map which have the key defined in the allowed keys.
func Pick[K comparable, V any](collection map[K]V, keys ...K) map[K]V {
	var result = make(map[K]V)

	for k := range collection {
		if Contains(keys, k) {
			result[k] = collection[k]
		}
	}

	return result
}

// PickBy extracts all the map elements for which the callback function returns truthy.
func PickBy[K comparable, V any](collection map[K]V, fn func(key K, val V) bool) map[K]V {
	var result = make(map[K]V)

	for k, v := range collection {
		if fn(k, v) {
			result[k] = collection[k]
		}
	}

	return result
}

// Omit is the opposite of Pick, it extracts all the map elements which keys are not omitted.
func Omit[K comparable, V any](collection map[K]V, keys ...K) map[K]V {
	for k := range collection {
		if Contains(keys, k) {
			delete(collection, k)
		}
	}

	return collection
}

// OmitBy is the opposite of Omit, it removes all the map elements for which the callback function returns true.
func OmitBy[K comparable, V any](collection map[K]V, fn func(key K, val V) bool) map[K]V {
	for k, v := range collection {
		if fn(k, v) {
			delete(collection, k)
		}
	}

	return collection
}

// PartitionMap split the collection into two arrays, the one whose elements satisfies the condition
// expressed in the callback function (fn) and one whose elements don't satisfies the condition.
func PartitionMap[K comparable, V any](mapSlice []map[K]V, fn func(map[K]V) bool) [2][]map[K]V {
	var result = [2][]map[K]V{}

	for _, m := range mapSlice {
		for k, v := range m {
			m[k] = v
			if fn(m) {
				result[0] = append(result[0], m)
				break
			} else {
				result[1] = append(result[1], m)
				break
			}
		}
	}

	return result
}
