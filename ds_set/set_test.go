package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New[int]()
	assert.NotNil(t, s)
	assert.NotNil(t, s.data)
	assert.Equal(t, 0, len(s.data))
}

func TestInsert(t *testing.T) {
	s := New[int]()

	s.Insert(10)
	assert.NotNil(t, s.data)
	assert.Equal(t, 1, len(s.data))
	_, ok := s.data[10]
	assert.True(t, ok)

	s.Insert(10)
	assert.Equal(t, 1, len(s.data))
	_, ok = s.data[10]
	assert.True(t, ok)

	s.Insert(20)
	assert.Equal(t, 2, len(s.data))
	_, ok = s.data[20]
	assert.True(t, ok)
}

func TestDelete(t *testing.T) {
	s := New[int]()

	s.Insert(10)
	s.Insert(20)
	s.Insert(30)

	s.Delete(10)
	_, ok := s.data[10]
	assert.False(t, ok)

	s.Delete(20)
	_, ok = s.data[20]
	assert.False(t, ok)

	s.Delete(30)
	_, ok = s.data[30]
	assert.False(t, ok)

	assert.Equal(t, 0, len(s.data))
}

func TestHas(t *testing.T) {
	s := New[int]()

	s.Insert(10)
	s.Insert(20)
	s.Insert(30)

	assert.True(t, s.Has(10))
	assert.True(t, s.Has(20))
	assert.True(t, s.Has(30))
	assert.False(t, s.Has(50))
}

func TestSize(t *testing.T) {
	s := New[int]()
	assert.Equal(t, 0, s.Size())

	for i := 1; i <= 5; i++ {
		s.Insert(i)
		assert.Equal(t, i, s.Size())
	}

	for i := 4; i <= 0; i++ {
		s.Delete(i)
		assert.Equal(t, i, s.Size())
	}
}
