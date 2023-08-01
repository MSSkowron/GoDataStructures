package set

type Set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable](values ...T) *Set[T] {
	set := &Set[T]{
		data: make(map[T]struct{}),
	}
	for _, v := range values {
		set.data[v] = struct{}{}
	}
	return set
}

func (s *Set[T]) Insert(value T) {
	s.data[value] = struct{}{}
}

func (s *Set[T]) Delete(value T) {
	delete(s.data, value)
}

func (s *Set[T]) Has(value T) bool {
	_, ok := s.data[value]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.data)
}
