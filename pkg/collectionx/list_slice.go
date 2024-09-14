package collectionx

import (
	"iter"
	"slices"
)

// SliceList is not thread-safe.
type SliceList[T any] struct {
	slice []T
}

var _ List[int] = (*SliceList[int])(nil)

func NewSliceList[T any](elements ...T) *SliceList[T] {
	slice := make([]T, len(elements))
	if len(elements) > 0 {
		copy(slice, elements)
	}

	return &SliceList[T]{
		slice: slice,
	}
}

func (l *SliceList[T]) Append(value T) {
	l.slice = append(l.slice, value)
}

func (l *SliceList[T]) At(index int) (T, bool) {
	if index < 0 || index >= len(l.slice) {
		var zero T
		return zero, false
	}
	return l.slice[index], true
}

func (l *SliceList[T]) UnsafeAt(index int) T {
	return l.slice[index]
}

func (l *SliceList[T]) Resize(size int) {
	ns := make([]T, size)
	if size == len(l.slice) {
		return
	}

	if size < len(l.slice) {
		copy(ns, l.slice[:size])
	} else {
		copy(ns, l.slice)
	}

	l.slice = ns
}

func (l *SliceList[T]) Len() int {
	return len(l.slice)
}

func (l *SliceList[T]) Elements() []T {
	return l.slice[:len(l.slice)]
}

func (l *SliceList[T]) Iterator() iter.Seq2[int, T] {
	return slices.All(l.slice)
}
