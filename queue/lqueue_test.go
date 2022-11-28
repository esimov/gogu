package queue

import (
	"fmt"
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
	q.Dequeue()

	q.Enqueue(10)
	assert.Equal(1, q.Size())
	q.Clear()
	assert.Equal(0, q.Size())
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

	item := q.Dequeue()
	assert.Equal(0, item)
	for q.Size() > 0 {
		item := q.Dequeue()
		assert.Equal(tmp[item], item)
	}
	assert.Equal(0, q.Size())
}

func Example_linkedQueue() {
	q := NewLinked(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	fmt.Println(q.Search(3))
	q.Dequeue()

	q.Enqueue(10)
	fmt.Println(q.Size())
	q.Clear()
	fmt.Println(q.Size())

	// Output:
	// 1
	// 2
	// 3
	// true
	// 1
	// 0
}
