package core

type Stack[T any] struct {
	items   []T
	maxSize int
}

func NewStack[T any](maxSize int) *Stack[T] {
	return &Stack[T]{
		items:   make([]T, 0),
		maxSize: maxSize,
	}
}

func (s *Stack[T]) Push(item T) {
	if len(s.items) < s.maxSize {
		s.items = append(s.items, item)
	}
}

func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue
	}

	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Items() []T {
	return s.items
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return item
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
