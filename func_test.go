package gogu

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
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
func TestFunc_Retry(t *testing.T) {
	assert := assert.New(t)

	n := -4
	idx := 0
	ForEach([]string{"one", "two", "three"}, func(val string) {
		rt := RType[string]{Input: val}
		attempts, e := rt.Retry(n, func(elem string) (err error) {
			if len(elem)%3 != 0 {
				err = fmt.Errorf("retry failed: number of %d attempts exceeded", n)
			}
			return err
		})
		assert.Error(e)
		assert.Equal(0, attempts)
	})

	n = 2
	ForEach([]string{"one", "two", "three"}, func(val string) {
		rt := RType[string]{Input: val}
		attempts, e := rt.Retry(n, func(elem string) (err error) {
			if len(elem)%3 != 0 {
				err = fmt.Errorf("retry failed: number of %d attempts exceeded", n)
			}
			return err
		})
		switch idx {
		case 0:
			assert.Equal(0, attempts)
			assert.NoError(e)
		case 1:
			assert.Equal(0, attempts)
			assert.NoError(e)
		case 2:
			assert.Equal(2, attempts)
			assert.Error(e)
		}
		idx++
	})
}

func TestFunc_RetryWithDelay(t *testing.T) {
	assert := assert.New(t)

	n := 5
	idx := 0
	ForEach([]int{1, 2, 3, 4, 5, 6, 7}, func(val int) {
		rt := RType[int]{Input: val}
		d, att, e := rt.RetryWithDelay(n, 20*time.Millisecond, func(d time.Duration, elem int) (err error) {
			if elem%2 != 0 {
				err = fmt.Errorf("retry failed: number of %d attempts exceeded", n)
			}
			return err
		})
		switch idx {
		case 0, 2:
			assert.GreaterOrEqual(int(d.Milliseconds()), 100)
			assert.Equal(5, att)
			assert.Error(e)
		case 1, 3:
			assert.LessOrEqual(int(d.Milliseconds()), 10)
			assert.Equal(0, att)
			assert.NoError(e)
		}
		idx++
	})

	// Here we are simulating an external service. In case the response time
	// exceeds a certain limit we are stop retrying and we are returning an error.
	services := []struct {
		service string
		time    int
	}{
		{
			service: "AWS1",
			time:    10,
		}, {
			service: "AWS2",
			time:    20,
		},
	}

	type Service[T ~string, N constraints.Integer] struct {
		Service T
		Time    N
	}

	for _, srv := range services {
		service := Service[string, int]{
			Service: srv.service,
			Time:    srv.time,
		}
		rtyp := RType[Service[string, int]]{
			Input: service,
		}

		d, att, e := rtyp.RetryWithDelay(n, 20*time.Millisecond, func(d time.Duration, srv Service[string, int]) (err error) {
			if srv.Time > 10 {
				err = fmt.Errorf("retry failed: service time exceeded")
			}
			return err
		})

		fmt.Println(d, att, e)
	}
}
