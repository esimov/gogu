package gogu

import "math/rand"

// Shuffle implements the Fisher-Yates shuffle algorithm applied to a slice.
func Shuffle[T any](collection []T) []T {
	for i := len(collection) - 1; i > 0; i-- {
		j := rand.Int() % (i + 1)
		swap(&collection[i], &collection[j])
	}
	return collection
}

// swap the two items.
func swap[T any](a, b *T) {
	tmp := *a
	*a = *b
	*b = tmp
}
