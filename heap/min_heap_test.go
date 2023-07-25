package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap_New(t *testing.T) {
	mh := NewMinHeap[int]()
	assert.NotNil(t, mh)
	assert.NotNil(t, mh.arr)
	assert.Equal(t, 0, len(mh.arr))
}

func TestMinHeap_Insert(t *testing.T) {
	mh := NewMinHeap[int]()

	mh.Insert(1)
	assert.Equal(t, 1, len(mh.arr))
	assert.Equal(t, 1, mh.arr[0])

	mh.Insert(2)
	assert.Equal(t, 2, len(mh.arr))
	assert.Equal(t, 1, mh.arr[0])
	assert.Equal(t, 2, mh.arr[1])

	mh.Insert(3)
	assert.Equal(t, 3, len(mh.arr))
	assert.Equal(t, 1, mh.arr[0])
	assert.Equal(t, 2, mh.arr[1])
	assert.Equal(t, 3, mh.arr[2])

	mh.Insert(0)
	assert.Equal(t, 4, len(mh.arr))
	assert.Equal(t, 0, mh.arr[0])
	assert.Equal(t, 1, mh.arr[1])
	assert.Equal(t, 3, mh.arr[2])
	assert.Equal(t, 2, mh.arr[3])
}

func TestMinHeap_Extract(t *testing.T) {
	mh := NewMinHeap[int]()
	mh.Insert(1)
	mh.Insert(2)
	mh.Insert(3)
	mh.Insert(0)

	min, err := mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 0, min)
	assert.Equal(t, 3, len(mh.arr))
	assert.Equal(t, 1, mh.arr[0])
	assert.Equal(t, 2, mh.arr[1])
	assert.Equal(t, 3, mh.arr[2])

	min, err = mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 1, min)
	assert.Equal(t, 2, len(mh.arr))
	assert.Equal(t, 2, mh.arr[0])
	assert.Equal(t, 3, mh.arr[1])

	min, err = mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 2, min)
	assert.Equal(t, 1, len(mh.arr))
	assert.Equal(t, 3, mh.arr[0])

	min, err = mh.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 3, min)
	assert.Equal(t, 0, len(mh.arr))

	_, err = mh.Extract()
	assert.NotNil(t, err)
	assert.Equal(t, ErrHeapIsEmpty, err)
}

func TestMinHeap_IsEmpty(t *testing.T) {
	mh := NewMinHeap[int]()
	assert.True(t, mh.IsEmpty())

	mh.Insert(1)
	assert.False(t, mh.IsEmpty())

	_, _ = mh.Extract()
	assert.True(t, mh.IsEmpty())
}

func TestMinHeap_Size(t *testing.T) {
	mh := NewMinHeap[int]()
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
