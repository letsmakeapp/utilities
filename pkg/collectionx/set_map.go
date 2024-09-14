package collectionx

type MapSet[T comparable] struct {
	m map[T]struct{}
}

var _ Set[int] = (*MapSet[int])(nil)

func NewMapSet[T comparable]() *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}),
	}
}

func (s *MapSet[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *MapSet[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *MapSet[T]) Contains(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *MapSet[T]) Elements() []T {
	elements := make([]T, 0, len(s.m))
	for k := range s.m {
		elements = append(elements, k)
	}
	return elements
}

func (s *MapSet[T]) Size() int {
	return len(s.m)
}

func (s *MapSet[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *MapSet[T]) IsNotEmpty() bool {
	return len(s.m) > 0
}
