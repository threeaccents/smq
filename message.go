package smq

//Message is
type Message struct {
	Payload []byte
	done    chan int
}

//Finish is
func (m Message) Finish() {
	<-m.done
}
