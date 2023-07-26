package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testingStruct struct {
	a int
	b string
}

func TestNew(t *testing.T) {
	pq := New[int, testingStruct]()

	assert.NotNil(t, pq)
	assert.NotNil(t, pq.arr)
	assert.Equal(t, 0, len(pq.arr))
}

func TestEnqueue(t *testing.T) {
	pq := New[int, testingStruct]()

	pq.Enqueue(1, testingStruct{
		a: 1,
		b: "a",
	})
	assert.Equal(t, 1, len(pq.arr))
	assert.Equal(t, 1, pq.arr[0].priority)
	assert.Equal(t, testingStruct{
		a: 1,
		b: "a",
	}, pq.arr[0].value)

	pq.Enqueue(2, testingStruct{
		a: 2,
		b: "b",
	})
	assert.Equal(t, 2, len(pq.arr))
	assert.Equal(t, 2, pq.arr[0].priority)
	assert.Equal(t, testingStruct{
		a: 2,
		b: "b",
	}, pq.arr[0].value)
	assert.Equal(t, 1, pq.arr[1].priority)
	assert.Equal(t, testingStruct{
		a: 1,
		b: "a",
	}, pq.arr[1].value)

	pq.Enqueue(3, testingStruct{
		a: 3,
		b: "c",
	})
	assert.Equal(t, 3, len(pq.arr))
	assert.Equal(t, 3, pq.arr[0].priority)
	assert.Equal(t, testingStruct{
		a: 3,
		b: "c",
	}, pq.arr[0].value)
	assert.Equal(t, 1, pq.arr[1].priority)
	assert.Equal(t, testingStruct{
		a: 1,
		b: "a",
	}, pq.arr[1].value)
	assert.Equal(t, 2, pq.arr[2].priority)
	assert.Equal(t, testingStruct{
		a: 2,
		b: "b",
	}, pq.arr[2].value)
}

func TestDequeue(t *testing.T) {
	pq := New[int, testingStruct]()
	pq.Enqueue(1, testingStruct{
		a: 1,
		b: "a",
	})
	pq.Enqueue(2, testingStruct{
		a: 2,
		b: "b",
	})
	pq.Enqueue(3, testingStruct{
		a: 3,
		b: "c",
	})

	value, err := pq.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, testingStruct{
		a: 3,
		b: "c",
	}, value)
	assert.Equal(t, 2, len(pq.arr))
	assert.Equal(t, 2, pq.arr[0].priority)
	assert.Equal(t, testingStruct{
		a: 2,
		b: "b",
	}, pq.arr[0].value)
	assert.Equal(t, 1, pq.arr[1].priority)
	assert.Equal(t, testingStruct{
		a: 1,
		b: "a",
	}, pq.arr[1].value)

	value, err = pq.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, testingStruct{
		a: 2,
		b: "b",
	}, value)
	assert.Equal(t, 1, len(pq.arr))
	assert.Equal(t, 1, pq.arr[0].priority)
	assert.Equal(t, testingStruct{
		a: 1,
		b: "a",
	}, pq.arr[0].value)

	value, err = pq.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, testingStruct{
		a: 1,
		b: "a",
	}, value)
	assert.Equal(t, 0, len(pq.arr))

	_, err = pq.Dequeue()
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriorityQueueIsEmpty, err)
}

func TestPeek(t *testing.T) {
	pq := New[int, testingStruct]()
	pq.Enqueue(1, testingStruct{
		a: 1,
		b: "a",
	})
	pq.Enqueue(2, testingStruct{
		a: 2,
		b: "b",
	})
	pq.Enqueue(3, testingStruct{
		a: 3,
		b: "c",
	})

	value, err := pq.Peek()
	assert.Nil(t, err)
	assert.Equal(t, testingStruct{
		a: 3,
		b: "c",
	}, value)

	_, _ = pq.Dequeue()

	value, err = pq.Peek()
	assert.Nil(t, err)
	assert.Equal(t, testingStruct{
		a: 2,
		b: "b",
	}, value)

	_, _ = pq.Dequeue()

	value, err = pq.Peek()
	assert.Nil(t, err)
	assert.Equal(t, testingStruct{
		a: 1,
		b: "a",
	}, value)

	_, _ = pq.Dequeue()

	_, err = pq.Peek()
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriorityQueueIsEmpty, err)
}

func TestIsEmpty(t *testing.T) {
	pq := New[int, testingStruct]()
	assert.True(t, pq.IsEmpty())

	pq.Enqueue(1, testingStruct{
		a: 1,
		b: "a",
	})
	assert.False(t, pq.IsEmpty())

	_, _ = pq.Dequeue()
	assert.True(t, pq.IsEmpty())
}

func TestSize(t *testing.T) {
	pq := New[int, testingStruct]()
	assert.Equal(t, 0, pq.Size())

	pq.Enqueue(1, testingStruct{
		a: 1,
		b: "a",
	})
	assert.Equal(t, 1, pq.Size())

	pq.Enqueue(2, testingStruct{
		a: 2,
		b: "b",
	})
	assert.Equal(t, 2, pq.Size())

	pq.Enqueue(3, testingStruct{
		a: 3,
		b: "c",
	})
	assert.Equal(t, 3, pq.Size())

	_, _ = pq.Dequeue()
	assert.Equal(t, 2, pq.Size())

	_, _ = pq.Dequeue()
	assert.Equal(t, 1, pq.Size())

	_, _ = pq.Dequeue()
	assert.Equal(t, 0, pq.Size())
}
