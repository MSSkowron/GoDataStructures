package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	trie := New()

	assert.NotNil(t, trie)
	assert.NotNil(t, trie.root)
	assert.False(t, trie.root.isEnd)

	for _, child := range trie.root.children {
		assert.Nil(t, child)
	}
}

func TestInsert(t *testing.T) {
	trie := New()

	trie.Insert("abc")
	assert.True(t, trie.root.children[0].children[1].children[2].isEnd)

	trie.Insert("")
	assert.True(t, trie.root.isEnd)
}

func TestDelete(t *testing.T) {
	trie := New()
	trie.Insert("abc")

	trie.Delete("abc")
	assert.False(t, trie.Search("abc"))

	trie.Delete("xyz")
	assert.False(t, trie.Search("xyz"))
}

func TestSearch(t *testing.T) {
	trie := New()
	trie.Insert("abc")

	assert.True(t, trie.Search("abc"))
	assert.False(t, trie.Search("a"))
	assert.False(t, trie.Search("ab"))
	assert.False(t, trie.Search("abcd"))
	assert.False(t, trie.Search("xyz"))
	assert.False(t, trie.Search(""))
}
