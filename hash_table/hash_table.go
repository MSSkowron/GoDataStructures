package hashtable

// ArrSize is the size of the hash table array
const ArrSize = 100

// HashTable is a data structure that stores key-value pairs.
// The hash table is implemented using separate chaining.
type HashTable struct {
	arr [ArrSize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

// New creates a new hash table
func New() *HashTable {
	ht := &HashTable{}
	ht.init()
	return ht
}

func (ht *HashTable) init() {
	for i := range ht.arr {
		ht.arr[i] = &bucket{}
	}
}

// Insert inserts a new key-value pair into the hash table
func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.arr[index].insert(key)
}

// Delete deletes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.arr[index].delete(key)
}

// Search searches for a key in the hash table
func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.arr[index].search(key)
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	}
}

func (b *bucket) delete(key string) {
	if b.head == nil {
		return
	}
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	currentNode := b.head
	for currentNode.next != nil {
		if currentNode.next.key == key {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArrSize
}
