package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)

	s := NewStack[int]()
	s.Push(1)
	assert.Equal(1, s.Peek())
	s.Push(2)
	assert.Equal(2, s.Peek())
	s.Pop()
	assert.Equal(1, s.Peek())
	assert.True(s.Search(1))
	s.Pop()
	assert.False(s.Search(1))
	assert.Empty(s.Size())

	s.Push(1)
	s.Pop()
	assert.False(s.Search(1))
}
