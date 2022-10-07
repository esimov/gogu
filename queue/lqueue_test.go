package queue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedQueue(t *testing.T) {
	assert := assert.New(t)

	q := NewLinked(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(1, q.Peek())
	q.Dequeue()
	assert.Equal(2, q.Peek())
	q.Dequeue()
	assert.Equal(3, q.Peek())
	assert.True(q.Search(3))
	_, err := q.Dequeue()
	assert.Error(err)
}

func TestLinkedQueue_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	q := NewLinked(0)

	n := 100
	tmp := make([]int, n)
	tmp[0] = 0

	for i := 1; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			q.Enqueue(i)

			mu.Lock()
			tmp[i] = i
			mu.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()
	assert.Equal(n, q.Size())
	assert.Equal(0, q.Peek())

	item, err := q.Dequeue()
	assert.NoError(err)
	assert.Equal(0, item)
	for {
		item, err := q.Dequeue()
		assert.Equal(tmp[item], item)
		if err != nil {
			break
		}
	}
	assert.Equal(tmp[n-1], q.Peek())
}
