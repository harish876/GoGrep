package lib

type Stack[T any] struct {
	container []T
}

func (s *Stack[T]) Push(val T) {
	s.container = append(s.container, val)
}

func (s *Stack[T]) Pop() bool {
	if s.Len() == 0 {
		return false
	}
	s.container = s.container[:len(s.container)-1]
	return true
}

// panics if stack is empty
func (s *Stack[T]) Top() T {
	return s.container[len(s.container)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.container)
}
