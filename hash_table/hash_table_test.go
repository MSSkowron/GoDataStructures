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
		assert.Nil(t, ht.arr[i].head)
	}
}

func TestInsert(t *testing.T) {
	ht := New()

	ht.Insert("foo")
	idx := hash("foo")
	assert.NotNil(t, ht.arr[idx].head)
	assert.Equal(t, "foo", ht.arr[idx].head.key)
	assert.Nil(t, ht.arr[idx].head.next)

	ht.Insert("bar")
	idx = hash("bar")
	assert.NotNil(t, ht.arr[idx].head)
	assert.Equal(t, "bar", ht.arr[idx].head.key)
	assert.Nil(t, ht.arr[idx].head.next)
}

func TestDelete(t *testing.T) {
	ht := New()
	ht.Insert("foo")
	ht.Insert("bar")

	ht.Delete("foo")
	idx := hash("foo")
	assert.Nil(t, ht.arr[idx].head)

	ht.Delete("bar")
	idx = hash("bar")
	assert.Nil(t, ht.arr[idx].head)
}

func TestSearch(t *testing.T) {
	ht := New()
	ht.Insert("foo")
	ht.Insert("bar")

	assert.True(t, ht.Search("foo"))
	assert.True(t, ht.Search("bar"))
	assert.False(t, ht.Search("baz"))
}
