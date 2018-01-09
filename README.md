# Simple Message Queue

A fun simple lightweight in-memory message queue.

## Install

```bash
go get github.com/rodzzlessa24/smq
```

Or here is a gist link: https://gist.github.com/rodzzlessa24/b6d6f77deb491d0179ed0f7b42893ce8

## Usage

```go
// first arguemnt takes the maxQueues. Second arguement takes the maxWorkers.
q := smq.New(1000, 50)

// This returns a channel that we can listen for "events"
msgs := q.Consume()

go func() {
	for msg := range msgs {
		go func(msg smq.Message) {
			fmt.Println("executing worker", string(msg.Payload))
			time.Sleep(2 * time.Second)
			msg.Finish()
		}(msg)
	}
}()

for i := 0; i < 2000; i++ {
	q.Push([]byte(fmt.Sprintf("my message %d", i)))
}
```