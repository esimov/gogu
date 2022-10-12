package torx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert := assert.New(t)
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(1, Min(values...))
	assert.Equal(10, Max(values...))
	assert.Equal(2, Abs(-2))

	assert.Equal(5, Clamp(10, 1, 5))
	assert.Equal(-2, Clamp(-4, -2, 10))
}
