package smq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testQueue := New(50)

	assert.NotNil(t, testQueue)
	assert.NotNil(t, testQueue.items)
	assert.NotNil(t, testQueue.mu)
	assert.Equal(t, testQueue.buffer, 50)
}
