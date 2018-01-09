package smq

//SMQ is the main controller for the queue.
type SMQ struct {
	jobQueue chan []byte
	consumer chan Message
	worker   chan int
}

//New takes in the max queue size and workers.
func New(maxQueue, maxWorkers int) *SMQ {
	q := &SMQ{
		jobQueue: make(chan []byte, maxQueue),
		consumer: make(chan Message),
		worker:   make(chan int, maxWorkers),
	}

	go q.listen()

	return q
}

//Push adds a payload to the queue
func (q *SMQ) Push(payload []byte) {
	q.jobQueue <- payload
}

//Consume will return a channel that new payloads from the queue will be sent to.
func (q *SMQ) Consume() <-chan Message {
	return q.consumer
}

func (q *SMQ) listen() {
	for payload := range q.jobQueue {
		q.worker <- 1
		q.consumer <- Message{
			done:    q.worker,
			Payload: payload,
		}
	}
}

//Message contains the payload and also has a channel to inform the queue the message is finished.
type Message struct {
	Payload []byte
	done    chan int
}

//Finish will let the queue know it are ready to take on another worker.
func (m Message) Finish() {
	<-m.done
}
