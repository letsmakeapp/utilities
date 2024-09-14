package collectionx

import "iter"

type List[T any] interface {
	Append(value T)
	// TODO: working with remove

	At(index int) (T, bool)
	UnsafeAt(index int) T

	Resize(size int)

	Len() int

	Elements() []T
	Iterator() iter.Seq2[int, T]
}
