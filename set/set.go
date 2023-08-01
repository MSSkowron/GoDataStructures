package set

// Set represents a set of values of type T.
type Set[T comparable] struct {
	data map[T]struct{}
}

// New creates a new, empty set of values of type T.
func New[T comparable](values ...T) *Set[T] {
	set := &Set[T]{
		data: make(map[T]struct{}),
	}
	for _, v := range values {
		set.data[v] = struct{}{}
	}
	return set
}

// Insert inserts a value into the set.
func (s *Set[T]) Insert(value T) {
	s.data[value] = struct{}{}
}

// Delete deletes a value from the set.
func (s *Set[T]) Delete(value T) {
	delete(s.data, value)
}

// Has returns true if the set contains the given value, false otherwise.
func (s *Set[T]) Has(value T) bool {
	_, ok := s.data[value]
	return ok
}

// Size returns the size of the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Union returns the union of two sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := New[T]()

	for v := range s.data {
		unionSet.Insert(v)
	}

	for v := range other.data {
		unionSet.Insert(v)
	}

	return unionSet
}

// Intersection returns the intersection of two sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersectionSet := New[T]()

	for v := range s.data {
		if other.Has(v) {
			intersectionSet.Insert(v)
		}
	}

	return intersectionSet
}

// Difference returns the difference of two sets.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	differenceSet := New[T]()

	for v := range s.data {
		if !other.Has(v) {
			differenceSet.Insert(v)
		}
	}

	return differenceSet
}
