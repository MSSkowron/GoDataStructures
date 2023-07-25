package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap_New(t *testing.T) {
	mh := NewMaxHeap[int]()
	assert.NotNil(t, mh)
	assert.NotNil(t, mh.arr)
	assert.Equal(t, 0, len(mh.arr))
}

func TestMaxHeap_Insert(t *testing.T) {
	mh := NewMaxHeap[int]()

	mh.Insert(1)
	assert.Equal(t, 1, len(mh.arr))
	assert.Equal(t, 1, mh.arr[0])

	mh.Insert(2)
	assert.Equal(t, 2, len(mh.arr))
	assert.Equal(t, 2, mh.arr[0])
	assert.Equal(t, 1, mh.arr[1])

	mh.Insert(3)
	assert.Equal(t, 3, len(mh.arr))
	assert.Equal(t, 3, mh.arr[0])
	assert.Equal(t, 1, mh.arr[1])
	assert.Equal(t, 2, mh.arr[2])

	mh.Insert(0)
	assert.Equal(t, 4, len(mh.arr))
	assert.Equal(t, 3, mh.arr[0])
	assert.Equal(t, 1, mh.arr[1])
	assert.Equal(t, 2, mh.arr[2])
	assert.Equal(t, 0, mh.arr[3])
}

func TestMaxHeap_Extract(t *testing.T) {
	mh := NewMaxHeap[int]()
	mh.Insert(1)
	mh.Insert(2)
	mh.Insert(3)
	mh.Insert(0)

	max, err := mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 3, max)
	assert.Equal(t, 3, len(mh.arr))
	assert.Equal(t, 2, mh.arr[0])
	assert.Equal(t, 1, mh.arr[1])
	assert.Equal(t, 0, mh.arr[2])

	max, err = mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 2, max)
	assert.Equal(t, 2, len(mh.arr))
	assert.Equal(t, 1, mh.arr[0])
	assert.Equal(t, 0, mh.arr[1])

	max, err = mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 1, max)
	assert.Equal(t, 1, len(mh.arr))
	assert.Equal(t, 0, mh.arr[0])

	max, err = mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 0, max)
	assert.Equal(t, 0, len(mh.arr))

	_, err = mh.Extract()
	assert.NotNil(t, err)
	assert.Equal(t, ErrHeapIsEmpty, err)
}

func TestMaxHeap_IsEmpty(t *testing.T) {
	mh := NewMaxHeap[int]()
	assert.True(t, mh.IsEmpty())

	mh.Insert(1)
	assert.False(t, mh.IsEmpty())

	_, _ = mh.Extract()
	assert.True(t, mh.IsEmpty())
}

func TestMaxHeap_Size(t *testing.T) {
	mh := NewMaxHeap[int]()
	assert.Equal(t, 0, mh.Size())

	mh.Insert(1)
	assert.Equal(t, 1, mh.Size())

	mh.Insert(2)
	assert.Equal(t, 2, mh.Size())

	_, _ = mh.Extract()
	assert.Equal(t, 1, mh.Size())

	_, _ = mh.Extract()
	assert.Equal(t, 0, mh.Size())
}
