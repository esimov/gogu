package gogu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	assert := assert.New(t)

	src := []int{1, 2, 3, 4, 5, 6}
	shuffled := Shuffle(src)

	assert.ElementsMatch(src, shuffled)
	assert.NotEqual(src, shuffled)
}
