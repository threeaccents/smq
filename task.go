package smq

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Task is
type Task struct {
	ID        string
	Body      []byte
	Processed bool
	Attempts  int
	Timestamp time.Time
}

// Produce creates a new task and sends it to the queue.
func (q *Queue) Produce(topic string, v interface{}) error {
	payload, err := Marshal(v)
	if err != nil {
		return fmt.Errorf("error marshaling payload %v", err)
	}

	q.mu.Lock()
	// get all the topic listeners
	listeners, ok := q.items[topic]
	if !ok {
		var newListeners []chan Task
		q.items[topic] = append(newListeners, make(chan Task, q.buffer))
		listeners = q.items[topic]
	}
	q.mu.Unlock()
	task := Task{
		ID:        uuid.NewV4().String(),
		Body:      payload,
		Timestamp: time.Now(),
	}

	go func() {
		for _, listener := range listeners {
			listener <- task
		}
	}()

	return nil
}
