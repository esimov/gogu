package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	assert := assert.New(t)

	list := Init(1)
	assert.Equal(1, list.data)
	// cannot delete the first node if there is only one item in the list.
	_, err := list.Pop()
	assert.Error(err)
	err = list.Delete(&list.singleNode)
	assert.Error(err)

	list.Append(2)
	err = list.Delete(&list.singleNode) // delete first node
	assert.NoError(err)
	assert.Equal(2, list.singleNode.data)

	list.Unshift(1)
	assert.Equal(1, list.singleNode.data)

	last := list.Append(3)
	err = list.InsertAfter(last, 4)
	assert.NoError(err)
	list.Append(5)

	last = list.Append(6)
	list.Append(8)
	list.InsertAfter(last, 7)
	last = list.Append(9)

	err = list.Delete(last)
	assert.NoError(err)

	i := 0
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	item, err := list.Pop()
	assert.NoError(err)
	assert.Equal(7, item.data)

	i = 0
	expected = []int{1, 2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	item = list.Shift()
	assert.Equal(1, item.data)

	i = 0
	expected = []int{2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	item, err = list.Replace(20, 10)
	assert.Error(err)
	assert.Nil(item)

	item, err = list.Replace(7, 8)
	assert.NoError(err)
	assert.Equal(8, item.data)

	list.Unshift(1)
	list.Replace(8, 7)

	i = 0
	expected = []int{1, 2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	item, found := list.Find(7)
	assert.Equal(7, item.data)
	assert.True(found)

	item, found = list.Find(22)
	assert.Nil(item)
	assert.False(found)
}
