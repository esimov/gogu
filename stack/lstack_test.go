package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedStack(t *testing.T) {
	assert := assert.New(t)

	s := NewLinked(1)
	assert.Equal(1, s.Peek())
	s.Push(2)
	assert.Equal(2, s.Peek())
	s.Pop()
	assert.Equal(1, s.Peek())
	assert.True(s.Search(1))
	_, err := s.Pop()
	assert.Error(err)

	s.Push(1)
	s.Push(2)
	_, err = s.Pop()
	assert.NoError(err)
	assert.True(s.Search(1))
}
