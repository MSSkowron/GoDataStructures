package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	q := New[int]()

	assert.NotNil(t, q)
	assert.Equal(t, 0, len(q.data))
}

func TestEnqueue(t *testing.T) {
	q := New[int]()

	q.Enqueue(10)

	assert.Equal(t, 1, len(q.data))
	assert.Equal(t, 10, q.data[0])

	q.Enqueue(20)

	assert.Equal(t, 2, len(q.data))
	assert.Equal(t, 20, q.data[1])

	q.Enqueue(30)

	assert.Equal(t, 3, len(q.data))
	assert.Equal(t, 30, q.data[2])
}

func TestDequeue(t *testing.T) {
	q := New[int]()

	for i := 10; i <= 100; i += 10 {
		q.Enqueue(i)
	}

	for i := 10; i <= 100; i += 10 {
		item, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, i, item)
	}

	item, err := q.Dequeue()
	assert.ErrorIs(t, ErrEmptyQueue, err)
	assert.Equal(t, 0, item)
}

func TestIsEmpty(t *testing.T) {
	q := New[int]()

	isEmpty := q.IsEmpty()
	assert.Equal(t, true, isEmpty)

	q.Enqueue(10)

	isEmpty = q.IsEmpty()
	assert.Equal(t, false, isEmpty)

	_, _ = q.Dequeue()

	isEmpty = q.IsEmpty()
	assert.Equal(t, true, isEmpty)
}
