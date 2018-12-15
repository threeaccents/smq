package smq

import (
	"testing"
)

var testConsumerFunc = func(task Task) {

}

type testConsumer struct {
}

func (c *testConsumer) HandleConsume(t Task) {}

func TestConsumer(t *testing.T) {
	testQueue := New(50)
	tests := []struct {
		topic    string
		consumer Consumer
		expected bool
	}{
		{"test", &testConsumer{}, true},
		{"", &testConsumer{}, false},
		{"test", nil, false},
	}

	for _, test := range tests {
		err := testQueue.Consumer(test.topic, 1, test.consumer)
		result := err == nil
		if result != test.expected {
			t.Errorf("expected %v got %v when error was %v", test.expected, result, err)
		}
	}
}

func TestConsumerIsAddedToQueue(t *testing.T) {
	testQueue := New(50)
	testTopic := "test_topic"

	initialLen := len(testQueue.items)

	testQueue.Consumer(testTopic, 1, &testConsumer{})

	itemsLen := len(testQueue.items)

	if initialLen == itemsLen {
		t.Error("Consumer not being added to queue")
	}

	if _, ok := testQueue.items[testTopic]; !ok {
		t.Error("Consumer topic not being added to queue")
	}
}

func TestMultipleConsumersForOneTopic(t *testing.T) {
	testQueue := New(50)
	testTopic := "test_topic"

	testQueue.Consumer(testTopic, 1, &testConsumer{})
	initialLen := len(testQueue.items[testTopic])
	testQueue.Consumer(testTopic, 1, &testConsumer{})
	consumersLen := len(testQueue.items[testTopic])

	if initialLen == consumersLen {
		t.Error("multiple consumers not being added to queue")
	}

}

func TestConsumerFunc(t *testing.T) {
	testQueue := New(50)
	tests := []struct {
		consumer func(task Task)
		expected bool
	}{
		{testConsumerFunc, true},
		{nil, false},
	}

	for _, test := range tests {
		err := testQueue.ConsumerFunc("test_topic", 1, test.consumer)
		result := err == nil
		if result != test.expected {
			t.Errorf("expected %v got %v when error was %v", test.expected, result, err)
		}
	}
}

func TestConsumerIsAddedToQueue_ConsumerFunc(t *testing.T) {
	testQueue := New(50)
	testTopic := "test_topic"

	initialLen := len(testQueue.items)

	testQueue.ConsumerFunc(testTopic, 1, testConsumerFunc)

	itemsLen := len(testQueue.items)

	if initialLen == itemsLen {
		t.Error("Consumer not being added to queue")
	}

	if _, ok := testQueue.items[testTopic]; !ok {
		t.Error("Consumer topic not being added to queue")
	}
}

func TestMultipleConsumersForOneTopic_ConsumerFunc(t *testing.T) {
	testQueue := New(50)
	testTopic := "test_topic"

	testQueue.ConsumerFunc(testTopic, 1, testConsumerFunc)
	initialLen := len(testQueue.items[testTopic])
	testQueue.ConsumerFunc(testTopic, 1, testConsumerFunc)
	consumersLen := len(testQueue.items[testTopic])

	if initialLen == consumersLen {
		t.Error("multiple consumers not being added to queue")
	}

}
