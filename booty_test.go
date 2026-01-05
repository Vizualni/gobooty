package gobooty

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestOneSingleCall(t *testing.T) {
	var count int32
	fn := One(func() int {
		return int(atomic.AddInt32(&count, 1))
	})

	if got := fn(); got != 1 {
		t.Fatalf("result=%d, want 1", got)
	}
	if got := atomic.LoadInt32(&count); got != 1 {
		t.Fatalf("factory called %d times, want 1", got)
	}
}

func TestOneSequentialCalls(t *testing.T) {
	var count int32
	fn := One(func() int {
		return int(atomic.AddInt32(&count, 1))
	})

	results := make([]int, 5)
	for i := range results {
		results[i] = fn()
	}

	for i, got := range results {
		if got != 1 {
			t.Fatalf("result[%d]=%d, want 1", i, got)
		}
	}
	if got := atomic.LoadInt32(&count); got != 1 {
		t.Fatalf("factory called %d times, want 1", got)
	}
}

func TestOneConcurrentCalls(t *testing.T) {
	var count int32
	fn := One(func() int {
		return int(atomic.AddInt32(&count, 1))
	})

	const calls = 20
	results := make([]int, calls)
	var wg sync.WaitGroup
	wg.Add(calls)
	for i := range calls {
		go func() {
			defer wg.Done()
			results[i] = fn()
		}()
	}
	wg.Wait()

	for i, got := range results {
		if got != 1 {
			t.Fatalf("result[%d]=%d, want 1", i, got)
		}
	}
	if got := atomic.LoadInt32(&count); got != 1 {
		t.Fatalf("factory called %d times, want 1", got)
	}
}

func TestTwoSingleCall(t *testing.T) {
	var count int32
	fn := Two(func() (int, string) {
		c := atomic.AddInt32(&count, 1)
		return int(c), fmt.Sprintf("v-%d", c)
	})

	a, b := fn()
	if a != 1 || b != "v-1" {
		t.Fatalf("result=(%d,%q), want (1,\"v-1\")", a, b)
	}
	if got := atomic.LoadInt32(&count); got != 1 {
		t.Fatalf("factory called %d times, want 1", got)
	}
}

func TestTwoSequentialCalls(t *testing.T) {
	var count int32
	fn := Two(func() (int, string) {
		c := atomic.AddInt32(&count, 1)
		return int(c), fmt.Sprintf("v-%d", c)
	})

	type pair struct {
		a int
		b string
	}

	results := make([]pair, 5)
	for i := range results {
		a, b := fn()
		results[i] = pair{a: a, b: b}
	}

	for i, got := range results {
		if got.a != 1 || got.b != "v-1" {
			t.Fatalf("result[%d]=(%d,%q), want (1,\"v-1\")", i, got.a, got.b)
		}
	}
	if got := atomic.LoadInt32(&count); got != 1 {
		t.Fatalf("factory called %d times, want 1", got)
	}
}

func TestTwoConcurrentCalls(t *testing.T) {
	var count int32
	fn := Two(func() (int, string) {
		c := atomic.AddInt32(&count, 1)
		return int(c), fmt.Sprintf("v-%d", c)
	})

	type pair struct {
		a int
		b string
	}

	const calls = 20
	results := make([]pair, calls)
	var wg sync.WaitGroup
	wg.Add(calls)
	for i := range calls {
		go func() {
			defer wg.Done()
			a, b := fn()
			results[i] = pair{a: a, b: b}
		}()
	}
	wg.Wait()

	for i, got := range results {
		if got.a != 1 || got.b != "v-1" {
			t.Fatalf("result[%d]=(%d,%q), want (1,\"v-1\")", i, got.a, got.b)
		}
	}
	if got := atomic.LoadInt32(&count); got != 1 {
		t.Fatalf("factory called %d times, want 1", got)
	}
}
