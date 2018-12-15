package smq

import "time"

// Task is
type Task struct {
	Body      []byte
	Processed bool
	Attempts  int
	Timestamp time.Time
}

// Produce creates a new task and sends it to the queue.
func (q *Queue) Produce(topic string, payload []byte) {
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
		Body:      payload,
		Timestamp: time.Now(),
	}

	for _, listner := range listeners {
		listner <- task
	}
}
