package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	tests := []struct {
		name    string
		arrange func(t *testing.T) *LinkedListStack[int]
		act     func(t *testing.T, stack *LinkedListStack[int])
		assert  func(t *testing.T, stack *LinkedListStack[int])
	}{
		{
			name: "empty",
			arrange: func(t *testing.T) *LinkedListStack[int] {
				return NewLinkedListStack[int]()
			},
			act: func(t *testing.T, stack *LinkedListStack[int]) {
				// DO NOTHING
			},
			assert: func(t *testing.T, stack *LinkedListStack[int]) {
				assert.True(t, stack.IsEmpty())
				assert.False(t, stack.IsNotEmpty())
				assert.Equal(t, 0, stack.Len())

				_, ok := stack.Peek()
				assert.False(t, ok)

				_, ok = stack.Pop()
				assert.False(t, ok)

				assert.Equal(t, []int{}, stack.Elements())
			},
		},
		{
			name: "push",
			arrange: func(t *testing.T) *LinkedListStack[int] {
				return NewLinkedListStack[int]()
			},
			act: func(t *testing.T, stack *LinkedListStack[int]) {
				stack.Push(1)
				stack.Push(2)
				stack.Push(3)
			},
			assert: func(t *testing.T, stack *LinkedListStack[int]) {
				assert.False(t, stack.IsEmpty())
				assert.True(t, stack.IsNotEmpty())
				assert.Equal(t, 3, stack.Len())

				v, ok := stack.Peek()
				assert.True(t, ok)
				assert.Equal(t, 3, v)

				assert.Equal(t, []int{3, 2, 1}, stack.Elements())
			},
		},
		{
			name: "push and pop",
			arrange: func(t *testing.T) *LinkedListStack[int] {
				return NewLinkedListStack[int]()
			},
			act: func(t *testing.T, stack *LinkedListStack[int]) {
				stack.Push(1)
			},
			assert: func(t *testing.T, stack *LinkedListStack[int]) {
				v, ok := stack.Pop()
				assert.True(t, ok)
				assert.Equal(t, 1, v)

				assert.True(t, stack.IsEmpty())
				assert.False(t, stack.IsNotEmpty())
				assert.Equal(t, 0, stack.Len())
				assert.Equal(t, []int{}, stack.Elements())
			},
		},
		{
			name: "push and peek",
			arrange: func(t *testing.T) *LinkedListStack[int] {
				return NewLinkedListStack[int]()
			},
			act: func(t *testing.T, stack *LinkedListStack[int]) {
				stack.Push(1)
			},
			assert: func(t *testing.T, stack *LinkedListStack[int]) {
				v, ok := stack.Peek()
				assert.True(t, ok)
				assert.Equal(t, 1, v)

				assert.False(t, stack.IsEmpty())
				assert.True(t, stack.IsNotEmpty())
				assert.Equal(t, 1, stack.Len())
				assert.Equal(t, []int{1}, stack.Elements())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.arrange(t)
			tt.act(t, stack)
			tt.assert(t, stack)
		})
	}
}
