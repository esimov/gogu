package queue

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	assert := assert.New(t)

	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(3, q.Size())
	assert.Equal(1, q.Peek())
	q.Dequeue()
	assert.Equal(2, q.Peek())
	q.Dequeue()
	assert.Equal(3, q.Peek())
	assert.True(q.Search(3))
	q.Dequeue()
	assert.Equal(0, q.Size())
	assert.False(q.Search(3))
}

func TestQueue_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}

	q := New[int]()
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			q.Enqueue(i)
			ch <- i
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for range ch {
		_, err := q.Dequeue()
		assert.NoError(err)
	}
	assert.Equal(0, q.Size())
}

func ExampleQueue() {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Size())
	fmt.Println(q.Peek())

	q.Dequeue()
	fmt.Println(q.Peek())
	fmt.Println(q.Search(2))

	// Output:
	// 3
	// 1
	// 2
	// true
}

func TestQueue_Race(t *testing.T) {
	const count = 10_000
	q := New[int]()
	for i := 0; i < 64; i++ {
		go func() {
			for i := 0; i < count; i++ {
				q.Peek()
			}
		}()
	}
	for i := 0; i < count; i++ {
		q.Enqueue(0)
		q.Dequeue()
	}
}
