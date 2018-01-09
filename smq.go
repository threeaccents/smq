package smq

//SMQ is
type SMQ struct {
	jobQueue chan []byte
	consumer chan Message
	worker   chan int
}

//New is
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

//Consume is
func (q *SMQ) Consume() <-chan Message {
	return q.consumer
}

func (q *SMQ) listen() {
	for {
		q.worker <- 1
		select {
		case payload := <-q.jobQueue:
			q.consumer <- Message{
				done:    q.worker,
				Payload: payload,
			}
		}
	}
}

//Message is
type Message struct {
	Payload []byte
	done    chan int
}

//Finish is
func (m Message) Finish() {
	<-m.done
}
