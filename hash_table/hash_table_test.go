package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ht := New[string]()
	assert.NotNil(t, ht)
	assert.Equal(t, ArrSize, len(ht.arr))

	for i := range ht.arr {
		assert.NotNil(t, ht.arr[i])
		assert.Equal(t, 0, ht.arr[i].Length())
	}
}

func TestInsert(t *testing.T) {
	htString := New[string]()

	htString.Insert("foo")
	idx := hash("foo")
	assert.Equal(t, 1, htString.arr[idx].Length())
	assert.True(t, htString.arr[idx].Contains("foo"))

	htString.Insert("bar")
	idx = hash("bar")
	assert.Equal(t, 1, htString.arr[idx].Length())
	assert.True(t, htString.arr[idx].Contains("bar"))

	htInt := New[int]()

	htInt.Insert(12)
	idx = hash(12)
	assert.Equal(t, 1, htInt.arr[idx].Length())
	assert.True(t, htInt.arr[idx].Contains(12))

	htInt.Insert(156)
	idx = hash(156)
	assert.Equal(t, 1, htInt.arr[idx].Length())
	assert.True(t, htInt.arr[idx].Contains(156))
}

func TestDelete(t *testing.T) {
	htString := New[string]()
	htString.Insert("foo")
	htString.Insert("bar")

	htString.Delete("foo")
	idx := hash("foo")
	assert.Equal(t, 0, htString.arr[idx].Length())
	assert.False(t, htString.arr[idx].Contains("foo"))

	htString.Delete("bar")
	idx = hash("bar")
	assert.Equal(t, 0, htString.arr[idx].Length())
	assert.False(t, htString.arr[idx].Contains("bar"))

	htInt := New[int]()
	htInt.Insert(12)
	htInt.Insert(156)

	htInt.Delete(12)
	idx = hash(12)
	assert.Equal(t, 0, htInt.arr[idx].Length())
	assert.False(t, htInt.arr[idx].Contains(12))

	htInt.Delete(156)
	idx = hash(156)
	assert.Equal(t, 0, htInt.arr[idx].Length())
	assert.False(t, htInt.arr[idx].Contains(156))
}

func TestSearch(t *testing.T) {
	htString := New[string]()
	htString.Insert("foo")
	htString.Insert("bar")

	assert.True(t, htString.Search("foo"))
	assert.True(t, htString.Search("bar"))
	assert.False(t, htString.Search("baz"))

	htInt := New[int]()
	htInt.Insert(12)
	htInt.Insert(156)

	assert.True(t, htInt.Search(12))
	assert.True(t, htInt.Search(156))
	assert.False(t, htInt.Search(243))
}
