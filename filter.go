package gogu

// Filter returns all the elements from the collection which satisfies the conditional logic of the callback function.
func Filter[T any](slice []T, fn func(T) bool) []T {
	res := make([]T, 0)

	for _, v := range slice {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

// Reject is the opposite of Filter.
// It returns the values from the collection without the elements for which the callback function returns true.
func Reject[T any](slice []T, fn func(val T) bool) []T {
	// TODO considering to create a new slice and append the values resulted
	// from the callback function, even if this imply a new allocation.
	for i := 0; i < len(slice); i++ {
		if fn(slice[i]) {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}

	return slice
}

// FilterMap iterates over the elements of a collection and returns a new collection
// representing all the items which satisfies the criteria formulated in the callback function.
func FilterMap[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V {
	filtered := map[K]V{}

	for k, v := range m {
		if fn(v) {
			filtered[k] = v
		}
	}

	return filtered
}

// FilterMapCollection filter out a one dimmensional collection of map items
// by applying the conditional logic of the callback function.
func FilterMapCollection[K comparable, V any](collection []map[K]V, fn func(V) bool) []map[K]V {
	filtered := []map[K]V{}

	for _, item := range collection {
		for _, v := range item {
			if fn(v) {
				filtered = append(filtered, item)
			}
		}
	}

	return filtered
}

// Filter2DMapCollection filter out a two dimmensional collection of map items
// by applying the conditional logic of the callback function.
func Filter2DMapCollection[K comparable, V any](collection []map[K]map[K]V, fn func(map[K]V) bool) []map[K]map[K]V {
	filtered := []map[K]map[K]V{}

	for _, item := range collection {
		for _, v := range item {
			if fn(v) {
				filtered = append(filtered, item)
			}
		}
	}

	return filtered
}
