package hashtable

import sll "github.com/MSSkowron/GoDataStructuresAndAlgorithms/linked_list/singly_linked_list"

// ArrSize is the size of the hash table array
const ArrSize = 100

// HashTable is a data structure that stores values based on theirs hash value.
// HashTable can be used with either int or string type.
type HashTable[T int | string] struct {
	arr [ArrSize]*sll.LinkedList[T]
}

// New creates a new hash table
func New[T int | string]() *HashTable[T] {
	ht := &HashTable[T]{}
	ht.init()
	return ht
}

func (ht *HashTable[T]) init() {
	for i := range ht.arr {
		ht.arr[i] = sll.New[T]()
	}
}

// Insert inserts a new value to the hash table
func (ht *HashTable[T]) Insert(value T) {
	index := hash(value)
	ht.arr[index].Prepend(value)
}

// Delete deletes a value from the hash table
func (ht *HashTable[T]) Delete(value T) {
	index := hash(value)
	ht.arr[index].Delete(value)
}

// Search searches for a value in the hash table
func (ht *HashTable[T]) Search(value T) bool {
	index := hash(value)
	return ht.arr[index].Contains(value)
}

func hash[T int | string](value T) int {
	switch any(value).(type) {
	case string:
		var sum int
		for _, v := range any(value).(string) {
			sum += int(v)
		}
		return sum % ArrSize
	case int:
		return any(value).(int) % ArrSize
	default:
		return -1
	}
}
