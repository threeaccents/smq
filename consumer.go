package smq

import "errors"

// Consumer handles an smq task.
type Consumer interface {
	HandleConsume(task Task)
}

// ConsumerFunc function type to inherit the Consumer interface
type ConsumerFunc func(task Task)

func (c ConsumerFunc) HandleConsume(task Task) {
	c(task)
}

// Consumer gets a task from the queue and handles the task.
func (q *Queue) Consumer(topic string, workers int, handler Consumer) error {
	if topic == "" {
		return errors.New("smq: nil topic")
	}
	if handler == nil {
		return errors.New("smq: nil handler")
	}
	listener := make(chan Task, q.buffer)
	q.mu.Lock()
	// topic listeners
	_, ok := q.items[topic]
	if !ok {
		// create new channel slice
		q.items[topic] = []chan Task{}
	}
	q.items[topic] = append(q.items[topic], listener)
	q.mu.Unlock()

	sem := make(chan int, workers)
	go func() {
		for task := range listener {
			sem <- 1
			go func(task Task) {
				handler.HandleConsume(task)
				<-sem
			}(task)
		}
	}()

	return nil
}

// ConsumerFunc helper function to handle a consumer. No signature needs to be created to match the Consumer interface.
func (q *Queue) ConsumerFunc(topic string, workers int, consumer func(task Task)) error {
	if consumer == nil {
		return errors.New("smq: nil consumer")
	}
	q.Consumer(topic, workers, ConsumerFunc(consumer))
	return nil
}
