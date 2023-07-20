package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New[int]()

	assert.NotNil(t, s)
	assert.Equal(t, 0, len(s.data))
}

func TestPush(t *testing.T) {
	s := New[int]()

	s.Push(10)

	assert.Equal(t, 1, len(s.data))
	assert.Equal(t, 10, s.data[0])

	s.Push(20)

	assert.Equal(t, 2, len(s.data))
	assert.Equal(t, 20, s.data[1])

	s.Push(30)

	assert.Equal(t, 3, len(s.data))
	assert.Equal(t, 30, s.data[2])
}

func TestPop(t *testing.T) {
	s := New[int]()

	for i := 10; i <= 100; i += 10 {
		s.Push(i)
	}

	for i := 100; i >= 10; i -= 10 {
		item, err := s.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, item)
	}

	item, err := s.Pop()
	assert.ErrorIs(t, ErrEmptyStack, err)
	assert.Equal(t, 0, item)
}

func TestPeek(t *testing.T) {
	s := New[int]()

	s.Push(10)

	item, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 10, item)

	s.Pop()

	item, err = s.Peek()
	assert.ErrorIs(t, ErrEmptyStack, err)
	assert.Equal(t, 0, item)
}

func TestSize(t *testing.T) {
	s := New[int]()

	for i := 10; i <= 100; i += 10 {
		s.Push(i)
	}

	size := s.Size()
	assert.Equal(t, 10, size)

	for i := 1; i <= 10; i++ {
		_, _ = s.Pop()

		size := s.Size()
		assert.Equal(t, 10-i, size)
	}

	size = s.Size()
	assert.Equal(t, 0, size)
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()

	s.Push(10)

	isEmpty := s.IsEmpty()
	assert.Equal(t, false, isEmpty)

	_, _ = s.Pop()

	isEmpty = s.IsEmpty()
	assert.Equal(t, true, isEmpty)
}
