package smq

import (
	"sync"
)

// Queue stores and synchronizes tasks among producers and consumers.
type Queue struct {
	mu     sync.Mutex
	items  map[string][]chan Task
	buffer int
}

// New creates and returns a new Queue.
func New(buffer int) *Queue {
	// buf := 50

	return &Queue{
		mu:     sync.Mutex{},
		items:  make(map[string][]chan Task),
		buffer: buffer,
	}
}
