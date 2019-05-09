# SMQ
A broadcasting in-memory message queue. You can provide how many workers you want per task and have multiple listeners per task.


## Usage
```go

type MyConsumer struct {
}

type payload struct {
	Msg string
}

func (c *MyConsumer) HandleConsume(t smq.Task) {
	var payload payload
	if err := smq.Unmarshal(t.Body, &payload); err != nil {
		log.Printf("unmarshaling task %s", err)
		return
	}
	fmt.Println("processing task", payload.Msg)
	time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
}

func main() {
	q := smq.New(1000)

	q.Consumer("test", 10, &MyConsumer{})

	for i := 0; i < 30; i++ {
		q.Produce("test", &payload{Msg: fmt.Sprintf("hello %d", i)})
	}

	fmt.Scanln()
}

```

Or using a ConsumerFunc helper method.

```go

func main() {
	q := smq.New(1000)

	q.ConsumerFunc("test", 10, func(t smq.Task) {
		var payload payload
		if err := smq.Unmarshal(t.Body, &payload); err != nil {
			log.Printf("unmarshaling task %s", err)
			return
		}
		fmt.Println("processing task", payload.Msg)
		time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	})

	for i := 0; i < 30; i++ {
		q.Produce("test", &payload{Msg: fmt.Sprintf("hello %d", i)})
	}

	fmt.Scanln()
}

```

With multiple listeners

```go
func main() {
	q := smq.New(1000)

	q.ConsumerFunc("test", 10, func(t smq.Task) {
		var payload payload
		if err := smq.Unmarshal(t.Body, &payload); err != nil {
			log.Printf("unmarshaling task %s", err)
			return
		}
		fmt.Println("processing task", payload.Msg)
		time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	})

	q.ConsumerFunc("test", 10, func(t smq.Task) {
		var payload payload
		if err := smq.Unmarshal(t.Body, &payload); err != nil {
			log.Printf("unmarshaling task %s", err)
			return
		}
		fmt.Println("processing task second handler", payload.Msg)
		time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	})

	for i := 0; i < 30; i++ {
		q.Produce("test", &payload{Msg: fmt.Sprintf("hello %d", i)})
	}

	fmt.Scanln()
}
```
