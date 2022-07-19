package gogu

import (
	"math/rand"
)

// Shuffle implements the Fisher-Yates shuffle algorithm applied to a slice.
func Shuffle[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)

	for i := len(src) - 1; i > 0; i-- {
		j := rand.Int() % (i + 1)
		swap(&dst[i], &dst[j])
	}
	return dst
}

// swap the two items.
func swap[T any](a, b *T) {
	tmp := *a
	*a = *b
	*b = tmp
}
