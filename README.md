# SMQ
A broadcasting in-memory message queue.

## Install

```bash
go get github.com/threeaccents/smq
```

## Usage
```go

type MyConsumer struct {
}

func (c *MyConsumer) HandleConsume(t smq.Task) {
	fmt.Println("processing task", string(t.Body))
	time.Sleep(3 * time.Second)
}

func main() {
	q := smq.New(1000)

	q.Consumer("test", 10, &MyConsumer{})

	for i := 0; i < 30; i++ {
		q.Produce("test", []byte(fmt.Sprintf("task number %d", i)))
	}

	fmt.Scanln()
}

```

Or using a ConsumerFunc helper method.

```go

func main() {
	q := smq.New(1000)

	q.ConsumerFunc("test", 10, func(task smq.Task) {
		fmt.Println("processing task", string(task.Body))
		time.Sleep(10 * time.Second)
	})

	for i := 0; i < 30; i++ {
		q.Produce("test", []byte(fmt.Sprintf("task number %d", i)))
	}

	fmt.Scanln()
}

```

With multiple listeners

```go
func main() {
	q := smq.New(1000)

	q.ConsumerFunc("test", 10, func(task smq.Task) {
		fmt.Println("processing task", string(task.Body))
		time.Sleep(10 * time.Second)
	})

	q.ConsumerFunc("test", 10, func(task smq.Task) {
		fmt.Println("processing task second handler", string(task.Body))
		time.Sleep(10 * time.Second)
	})

	for i := 0; i < 30; i++ {
		q.Produce("test", []byte(fmt.Sprintf("task number %d", i)))
	}

	fmt.Scanln()
}
```
