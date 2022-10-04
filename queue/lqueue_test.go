package queue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedQueue(t *testing.T) {
	assert := assert.New(t)

	ls := NewLinked(1)
	ls.Enqueue(2)
	ls.Enqueue(3)
	assert.Equal(1, ls.Peek())
	ls.Dequeue()
	assert.Equal(2, ls.Peek())
	ls.Dequeue()
	assert.Equal(3, ls.Peek())
	assert.True(ls.Search(3))
	_, err := ls.Dequeue()
	assert.Error(err)
}

func TestLinkedQueue_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	n := 10
	ls := NewLinked(0)
	tmp := make([]int, n)
	tmp[0] = 0

	for i := 1; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			ls.Enqueue(i)

			mu.Lock()
			tmp[i] = i
			mu.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()

	assert.Equal(0, ls.Peek())

	item, err := ls.Dequeue()
	assert.NoError(err)
	assert.Equal(0, item)
	for {
		item, err := ls.Dequeue()
		assert.Equal(tmp[item], item)
		if err != nil {
			break
		}
	}
	assert.Equal(tmp[n-1], ls.Peek())
}
