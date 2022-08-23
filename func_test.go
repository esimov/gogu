package gogu

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFunc_Flip(t *testing.T) {
	assert := assert.New(t)

	flipped := Flip(func(args ...int) []int {
		return ToSlice(args...)
	})
	assert.Equal([]int{}, flipped())
	assert.Equal([]int{1}, flipped(1))
	assert.Equal([]int{2, 1}, flipped(1, 2))
	assert.Equal([]int{3, 2, 1}, flipped(1, 2, 3))
}

func TestFunc_Delay(t *testing.T) {
	var (
		sample int
		mu     sync.Mutex
	)
	assert := assert.New(t)

	ch := make(chan struct{})
	now := time.Now()

	timer := Delay(20*time.Millisecond, func() {
		mu.Lock()
		sample = 1
		mu.Unlock()
		ch <- struct{}{}
	})
	mu.Lock()
	assert.Equal(0, sample)
	mu.Unlock()
	<-ch
	if timer.Stop() {
		<-timer.C
	}
	assert.Equal(1, sample)
	after := time.Since(now).Milliseconds()
	assert.LessOrEqual(int(after), 30)
	fmt.Println()
}

func TestFunc_After(t *testing.T) {
	assert := assert.New(t)

	sample := []int{1, 2, 3, 4, 5, 6}
	length := len(sample) - 1

	initVal := 0
	cb := func(val int) int {
		return val + 1
	}

	ForEach(sample, func(val int) {
		now := time.Now()
		After(&length, func() {
			<-time.After(10 * time.Millisecond)
			initVal = cb(initVal)
			after := time.Since(now).Milliseconds()
			assert.LessOrEqual(int(after), 20)
		})
	})
	assert.Equal(-1, length)
	assert.Equal(1, initVal)
}

func TestFunc_Before(t *testing.T) {
	assert := assert.New(t)
	c := NewCache[string, int](DefaultExpiration, NoExpiration)

	var n = 3
	sample := []int{1, 2, 3, 4, 5, 6}
	ForEach(sample, func(val int) {
		fn := func() int {
			<-time.After(10 * time.Millisecond)
			return n
		}
		res := Before(&n, c, fn)
		// The trick to test this function is to decrease the n value after each iteration.
		// We can be sure that the callback function is not served from the cache if n > 0.
		// In this case the cache item "func" should be empty.
		if n > 0 {
			val, _ := c.Get("func")
			assert.Nil(val)
			assert.Equal(res, n)
		}
		if n <= 0 {
			// Here we can be sure that the callback function is served from the cache.
			val, _ := c.Get("func")
			assert.NotNil(val)
			assert.Equal(res, 0)
		}
	})
}

func TestFunc_Once(t *testing.T) {
	assert := assert.New(t)
	c := NewCache[string, int](DefaultExpiration, NoExpiration)

	var n int = 2
	sample := []int{1, 2, 3, 4, 5, 6}
	ForEach(sample, func(val int) {
		fn := func() int {
			<-time.After(10 * time.Millisecond)
			n--
			return n
		}
		res := Once(c, fn)
		// If the callback function is invoked once the cache item "func" should be empty.
		// Otherwise the callback function is served from the cache.
		if n == 1 {
			val, _ := c.Get("func")
			assert.Nil(val)
		} else {
			val, _ := c.Get("func")
			assert.NotNil(val)
			assert.Equal(res, 0)
		}
	})
}
