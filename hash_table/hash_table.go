package hashtable

import sll "github.com/MSSkowron/GoDataStructuresAndAlgorithms/linked_list/singly_linked_list"

// ArrSize is the size of the hash table array
const ArrSize = 100

// HashTable is a data structure that stores key-value pairs.
// The hash table is implemented using separate chaining.
type HashTable struct {
	arr [ArrSize]*sll.LinkedList[string]
}

// New creates a new hash table
func New() *HashTable {
	ht := &HashTable{}
	ht.init()
	return ht
}

func (ht *HashTable) init() {
	for i := range ht.arr {
		ht.arr[i] = sll.New[string]()
	}
}

// Insert inserts a new key-value pair into the hash table
func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.arr[index].Prepend(key)
}

// Delete deletes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.arr[index].Delete(key)
}

// Search searches for a key in the hash table
func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.arr[index].Contains(key)
}

func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArrSize
}
