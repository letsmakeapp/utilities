package collectionx

type Set[T comparable] interface {
	Add(value T)
	Remove(value T)
	Contains(value T) bool
	Elements() []T
	Size() int
	IsEmpty() bool
	IsNotEmpty() bool
}
