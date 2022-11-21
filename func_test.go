package torx

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/esimov/torx/cache"
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

func Example_FuncFlip() {
	flipped := Flip(func(args ...int) []int {
		return ToSlice(args...)
	})
	fmt.Println(flipped(1, 2, 3))

	// Output:
	// [3 2 1]
}

func TestFunc_Delay(t *testing.T) {
	assert := assert.New(t)

	ch := make(chan struct{})
	now := time.Now()

	var value uint32
	timer := Delay(20*time.Millisecond, func() {
		atomic.AddUint32(&value, 1)
		ch <- struct{}{}
	})
	r1 := atomic.LoadUint32(&value)
	assert.Equal(0, int(r1))
	<-ch
	if timer.Stop() {
		<-timer.C
	}
	r1 = atomic.LoadUint32(&value)
	assert.Equal(1, int(r1))
	after := time.Since(now).Milliseconds()
	assert.LessOrEqual(int(after), 30)
}

func Example_FuncDelay() {
	ch := make(chan struct{})
	now := time.Now()

	var value uint32
	timer := Delay(20*time.Millisecond, func() {
		atomic.AddUint32(&value, 1)
		ch <- struct{}{}
	})
	r1 := atomic.LoadUint32(&value)
	fmt.Println(r1)
	<-ch
	if timer.Stop() {
		<-timer.C
	}
	r1 = atomic.LoadUint32(&value)
	fmt.Println(r1)
	after := time.Since(now).Milliseconds()
	fmt.Println(after)

	// Output:
	// 0
	// 1
	// 20
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

func Example_FuncAfter() {
	sample := []int{1, 2, 3, 4, 5, 6}
	length := len(sample) - 1

	initVal := 0
	fn := func(val int) int {
		return val + 1
	}

	ForEach(sample, func(val int) {
		now := time.Now()
		After(&length, func() {
			<-time.After(10 * time.Millisecond)
			initVal = fn(initVal)
			after := time.Since(now).Milliseconds()
			fmt.Println(after)
		})
	})

	// Output:
	// 10
}

func TestFunc_Before(t *testing.T) {
	assert := assert.New(t)
	c := cache.New[string, int](cache.DefaultExpiration, cache.NoExpiration)

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
			// Here the callback function is served from the cache.
			val, _ := c.Get("func")
			assert.NotNil(val)
			assert.Equal(res, 0)
		}
	})
}

func Example_FuncBefore() {
	c := cache.New[string, int](cache.DefaultExpiration, cache.NoExpiration)

	var n = 3
	sample := []int{1, 2, 3}
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
			fmt.Println(val)
			fmt.Println(res)
		}
		if n <= 0 {
			// Here the callback function is served from the cache.
			val, _ := c.Get("func")
			fmt.Println(val)
			fmt.Println(res)
		}
	})

	// Output:
	// <nil>
	// 2
	// <nil>
	// 1
	// &{0 0}
	// 0
}

func TestFunc_Once(t *testing.T) {
	assert := assert.New(t)
	c := cache.New[string, int](cache.DefaultExpiration, cache.NoExpiration)

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
	rand.Seed(time.Now().UTC().UnixNano())

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
	// exceeds a certain limit we stop retrying and we are returning an error.
	services := []struct {
		service string
		time    time.Duration
	}{
		{service: "AWS1"},
		{service: "AWS2"},
	}

	type Service[T ~string] struct {
		Service T
		Time    time.Duration
	}

	for _, srv := range services {
		r := random(1, 10)

		// Here we are simulating the response time of the external service
		// by generating some random duration between 1ms and 10ms.
		// All the test should pass because all of the responses are inside the predefined limit (10ms).
		service := Service[string]{
			Service: srv.service,
			Time:    time.Duration(r) * time.Millisecond,
		}
		rtyp := RType[Service[string]]{
			Input: service,
		}

		d, att, e := rtyp.RetryWithDelay(n, 20*time.Millisecond, func(d time.Duration, srv Service[string]) (err error) {
			if srv.Time.Milliseconds() > 10 {
				err = fmt.Errorf("retry failed: service time exceeded")
			}
			return err
		})
		assert.NoError(e)
		assert.Equal(0, att)
		assert.LessOrEqual(int(d.Milliseconds()), 10)
	}

	for _, srv := range services {
		r := random(20, 30)

		// Here we are simulating the response time of the external service
		// by generating some random duration between 20ms and 30ms.
		// All the test should fail because all of the responses are outside the predefined limit (10ms).
		service := Service[string]{
			Service: srv.service,
			Time:    time.Duration(r) * time.Millisecond,
		}
		rtyp := RType[Service[string]]{
			Input: service,
		}

		d, att, e := rtyp.RetryWithDelay(n, 20*time.Millisecond, func(d time.Duration, srv Service[string]) (err error) {
			if srv.Time.Milliseconds() > 10 {
				err = fmt.Errorf("retry failed: service time exceeded")
			}
			return err
		})
		assert.Error(e)
		assert.Equal(5, att)
		assert.Greater(int(d.Milliseconds()), 10)
	}
}

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func TestFunc_Debounce(t *testing.T) {
	assert := assert.New(t)

	var (
		counter1 uint64
		counter2 uint64
		counter3 uint64
	)

	f1 := func() {
		atomic.AddUint64(&counter1, 1)
	}

	f2 := func() {
		atomic.AddUint64(&counter2, 1)
	}

	f3 := func() {
		atomic.AddUint64(&counter3, 1)
	}

	debounce, cancel := NewDebounce(10 * time.Millisecond)
	for i := 0; i < 2; i++ {
		for j := 0; j < 100; j++ {
			debounce(f1)
		}
		<-time.After(20 * time.Millisecond)
	}
	cancel()

	debounce, cancel = NewDebounce(10 * time.Millisecond)
	for i := 0; i < 5; i++ {
		for j := 0; j < 50; j++ {
			debounce(f2)
		}
		for j := 0; j < 50; j++ {
			debounce(f2)
		}
		<-time.After(20 * time.Millisecond)
	}
	cancel()

	c1 := atomic.LoadUint64(&counter1)
	c2 := atomic.LoadUint64(&counter2)
	assert.Equal(2, int(c1))
	assert.Equal(5, int(c2))

	debounce, _ = NewDebounce(10 * time.Millisecond)
	var wg sync.WaitGroup

	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			debounce(f3)
			wg.Done()
		}()
	}
	wg.Wait()
	<-time.After(20 * time.Millisecond)

	c3 := atomic.LoadUint64(&counter3)
	assert.Equal(1, int(c3))

	debounce, cancel = NewDebounce(10 * time.Millisecond)
	atomic.SwapUint64(&counter3, 0)

	for i := 0; i < 2; i++ {
		for j := 0; j < 50; j++ {
			debounce(func() {
				atomic.AddUint64(&counter3, 1)
			})
		}
		if i == 1 {
			cancel()
		}
		<-time.After(20 * time.Millisecond)
	}

	c3 = atomic.LoadUint64(&counter3)
	assert.Equal(1, int(c3))
}

func Example_FuncDebounce() {
	var (
		counter1 uint64
		counter2 uint64
	)

	f1 := func() {
		atomic.AddUint64(&counter1, 1)
	}

	f2 := func() {
		atomic.AddUint64(&counter2, 1)
	}

	debounce, cancel := NewDebounce(10 * time.Millisecond)
	for i := 0; i < 2; i++ {
		for j := 0; j < 100; j++ {
			debounce(f1)
		}
		<-time.After(20 * time.Millisecond)
	}
	cancel()

	debounce, cancel = NewDebounce(10 * time.Millisecond)
	for i := 0; i < 5; i++ {
		for j := 0; j < 50; j++ {
			debounce(f2)
		}
		for j := 0; j < 50; j++ {
			debounce(f2)
		}
		<-time.After(20 * time.Millisecond)
	}
	cancel()

	c1 := atomic.LoadUint64(&counter1)
	c2 := atomic.LoadUint64(&counter2)
	fmt.Println(c1)
	fmt.Println(c2)

	// Output:
	// 2
	// 5
}

func TestFunc_ThrottleCache(t *testing.T) {
	var counter uint32
	var wg sync.WaitGroup

	assert := assert.New(t)

	c := cache.New[string, int](cache.DefaultExpiration, cache.NoExpiration)
	c.SetDefault("item", 0)

	limit := 100 * time.Millisecond
	throttle := NewThrottle(limit, true)

	wg.Add(1)
	go func() {
		for throttle.Next() {
			atomic.AddUint32(&counter, 1)
			ct := atomic.LoadUint32(&counter)
			c.Update("item", int(ct), cache.DefaultExpiration)
		}
		wg.Done()
	}()

	// This function should be called only once.
	for i := 0; i < 10; i++ {
		throttle.Call()
	}

	time.Sleep(limit)
	throttle.Cancel()

	wg.Wait()

	item, _ := c.Get("item")
	assert.Equal(1, item.Val())
}

func TestFunc_ThrottlePings(t *testing.T) {
	var wg sync.WaitGroup

	assert := assert.New(t)
	throttle := NewThrottle(time.Millisecond, true)

	count := 0
	wg.Add(1)
	go func() {
		for throttle.Next() {
			count += 1
		}
		wg.Done()
	}()

	time.Sleep(2 * time.Millisecond)
	throttle.Cancel()

	wg.Wait()

	assert.Equal(0, count)
}
