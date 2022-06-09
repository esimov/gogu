package gogu

func Filter[T any](s []T, fn func(T) bool) []T {
	rs := make([]T, 0)

	for _, v := range s {
		if fn(v) {
			rs = append(rs, v)
		}
	}

	return rs
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

// Filter2DMap is like FilterMap only applied on a two dimmensional map.
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
func Filter2DMapSlice[K comparable, V any](sm []map[K]map[K]V, fn func(map[K]V) bool) []map[K]map[K]V {
	filtered := []map[K]map[K]V{}

	for _, s := range sm {
		for _, v := range s {
			if fn(v) {
				filtered = append(filtered, s)
			}
		}
	}

	return filtered
}
