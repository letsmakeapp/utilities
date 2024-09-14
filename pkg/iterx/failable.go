package iterx

import (
	"errors"
	"fmt"
)

var (
	ErrUnwrapFailed = errors.New("unwrap failed")
)

type Failable[T any] struct {
	value T
	err   error
}

func NewFailableOk[T any](value T) Failable[T] {
	return Failable[T]{value: value}
}

func NewFailableErr[T any](err error) Failable[T] {
	return Failable[T]{err: err}
}

func (f Failable[T]) TryUnwrap() (T, error) {
	return f.value, f.err
}

func (f Failable[T]) Unwrap() T {
	if f.err != nil {
		panic(fmt.Errorf("%w: %v", ErrUnwrapFailed, f.err))
	}
	return f.value
}
