package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ht := New()
	assert.NotNil(t, ht)
	assert.Equal(t, ArrSize, len(ht.arr))

	for i := range ht.arr {
		assert.NotNil(t, ht.arr[i])
		assert.Equal(t, 0, ht.arr[i].Length())
	}
}

func TestInsert(t *testing.T) {
	ht := New()

	ht.Insert("foo")
	idx := hash("foo")
	assert.Equal(t, 1, ht.arr[idx].Length())
	assert.True(t, ht.arr[idx].Contains("foo"))

	ht.Insert("bar")
	idx = hash("bar")
	assert.Equal(t, 1, ht.arr[idx].Length())
	assert.True(t, ht.arr[idx].Contains("bar"))
}

func TestDelete(t *testing.T) {
	ht := New()
	ht.Insert("foo")
	ht.Insert("bar")

	ht.Delete("foo")
	idx := hash("foo")
	assert.Equal(t, 0, ht.arr[idx].Length())
	assert.False(t, ht.arr[idx].Contains("foo"))

	ht.Delete("bar")
	idx = hash("bar")
	assert.Equal(t, 0, ht.arr[idx].Length())
	assert.False(t, ht.arr[idx].Contains("bar"))
}

func TestSearch(t *testing.T) {
	ht := New()
	ht.Insert("foo")
	ht.Insert("bar")

	assert.True(t, ht.Search("foo"))
	assert.True(t, ht.Search("bar"))
	assert.False(t, ht.Search("baz"))
}
