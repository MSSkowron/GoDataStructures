package trie

// AlphabetSize is the number of possible characters in the trie.
const AlphabetSize = 26

type trieNode struct {
	children [AlphabetSize]*trieNode
	isEnd    bool
}

// Trie represents a trie data structure.
type Trie struct {
	root *trieNode
}

// New creates a new, empty trie.
func New() *Trie {
	return &Trie{
		root: &trieNode{},
	}
}

// Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
	curr := t.root
	for _, char := range word {
		charIndex := char - 'a'
		if curr.children[charIndex] == nil {
			curr.children[charIndex] = &trieNode{}
		}
		curr = curr.children[charIndex]
	}

	curr.isEnd = true
}

// Delete deletes a word from the trie.
func (t *Trie) Delete(word string) {
	t.delete(t.root, word, 0)
}

func (t *Trie) delete(curr *trieNode, word string, index int) {
	if index == len(word) {
		if curr.isEnd {
			curr.isEnd = false
		}
		return
	}

	charIndex := rune(word[index]) - 'a'
	if nextNode := curr.children[charIndex]; nextNode != nil {
		t.delete(nextNode, word, index+1)
	}
}

// Search returns true if the trie contains the given word, false otherwise.
func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, char := range word {
		charIndex := char - 'a'
		if curr.children[charIndex] == nil {
			return false
		}
		curr = curr.children[charIndex]
	}

	return curr.isEnd
}
