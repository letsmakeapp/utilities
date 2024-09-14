package stack

type Stack[T any] interface {
	Peek() (T, bool)
	Pop() (T, bool)
	Push(value T)
	Len() int
	IsEmpty() bool
	IsNotEmpty() bool
	Elements() []T
}
