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
