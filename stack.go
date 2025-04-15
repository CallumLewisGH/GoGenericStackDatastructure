package stack

import (
	"errors"
)

// Stack acts as a generic FILO data structure
type Stack[T any] struct {
	Data []T
}

// Push adds an element to the top of the stack
// Parameters:
//   - value is the value added to the top of the stack
//
// Returns:
//   - no returns
//
// Example
// stack{}.Push(4)
func (s *Stack[T]) Push(value T) {
	s.Data = append(s.Data, value)
}

// Pop returns the element at the top of the stack and removes it
// Parameters:
//   - no parameters
//
// Returns:
//   - returns the value at the top of the stack of type T
//
// Example
// stack{}.Pop()
func (s *Stack[T]) Pop() (T, error) {
	if len(s.Data) == 0 {
		var zero T
		return zero, errors.New("can't pop from empty stack")
	}
	var value T = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return value, nil
}

// Peak returns the element at the top of the stack and does NOT remove it
// Parameters:
//   - no parameters
//
// Returns:
//   - returns the value at the top of the stack of type T
//
// Example
// stack{}.Peek()
func (s *Stack[T]) Peek() (T, error) {
	if len(s.Data) == 0 {
		var zero T
		return zero, errors.New("can't peek an empty stack")
	}
	return s.Data[len(s.Data)-1], nil
}

// IsEmpty is used to determine if the stack is empty or not
// Parameters:
//   - no parameters
//
// Returns:
//   - returns boolean true or false if the stack is empty
//
// Example
// stack{}.Peek()
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Data) == 0
}
