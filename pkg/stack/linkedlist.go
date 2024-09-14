package stack

type node[T any] struct {
	value T
	next  *node[T]
}

type LinkedListStack[T any] struct {
	size int
	top  *node[T]
}

var _ Stack[int] = (*LinkedListStack[int])(nil)

func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		top: nil,
	}
}

func (l *LinkedListStack[T]) Elements() []T {
	elements := make([]T, 0, l.size)
	current := l.top
	for current != nil {
		elements = append(elements, current.value)
		current = current.next
	}
	return elements
}

func (l *LinkedListStack[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListStack[T]) IsNotEmpty() bool {
	return l.size > 0
}

func (l *LinkedListStack[T]) Len() int {
	return l.size
}

func (l *LinkedListStack[T]) Peek() (T, bool) {
	if l.size > 0 {
		return l.top.value, true
	}
	var zero T
	return zero, false
}

func (l *LinkedListStack[T]) Pop() (T, bool) {
	if l.size > 0 {
		top := l.top
		l.top = top.next
		l.size--
		return top.value, true
	}
	var zero T
	return zero, false
}

func (l *LinkedListStack[T]) Push(value T) {
	top := l.top
	node := &node[T]{value: value, next: top}
	l.top = node
	l.size++
}
