# Simple Message Queue

A fun simple lightweight in-memory message queue.

# Usage

```go
q := smq.New(2, 2)

msgs := q.Consume()

go func() {
	for msg := range msgs {
		go func(msg smq.Message) {
			fmt.Println("executing worker", string(msg.Payload))
			time.Sleep(5 * time.Second)
			msg.Finish()
		}(msg)
	}
}()

for i := 0; i < 200; i++ {
	q.Push([]byte(fmt.Sprintf("my message %d", i)))
}
```