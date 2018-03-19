package smq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testWorkers = 2
	testQLimit  = 1
)

func TestNew(t *testing.T) {
	q := New(testQLimit, testWorkers)
	assert.NotNil(t, q)
	assert.Equal(t, cap(q.worker), testWorkers)
	assert.Equal(t, cap(q.jobQueue), testQLimit)
}
