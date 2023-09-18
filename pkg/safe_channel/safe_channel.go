package safechannel

import (
	"github.com/samber/lo"
	"sync"
)

// SafeChannel is a concurrency-safe channel with capacity tracking.
type SafeChannel[T any] struct {
	ch  chan T
	len int
	cap int
	mu  sync.Mutex
}

// New creates a new SafeChannel with the given capacity.
func New[T any](capacity int) *SafeChannel[T] {
	return &SafeChannel[T]{
		ch:  make(chan T, capacity),
		cap: capacity,
	}
}

// Send safely sends a value to the channel, returning false if the channel is full.
func (sc *SafeChannel[T]) Send(value T) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if sc.len >= sc.cap {
		return false // Channel is full
	}

	sc.ch <- value
	sc.len++
	return true
}

// Receive safely receives a value from the channel, returning false if the channel is empty.
func (sc *SafeChannel[T]) Receive() (T, bool) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if sc.len <= 0 {
		return lo.Empty[T](), false // Channel is empty
	}

	value := <-sc.ch
	sc.len--
	return value, true
}
