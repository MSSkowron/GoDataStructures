package set

type Set[T comparable] map[T]struct{}

func New[T comparable](values ...T) *Set[T] {
	set := make(Set[T], 0)
	for _, v := range values {
		set[v] = struct{}{}
	}
	return &set
}

func (s Set[T]) Insert(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Delete(value T) {
	delete(s, value)
}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}
