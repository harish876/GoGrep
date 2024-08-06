package lib

import "testing"

func TestStack(t *testing.T) {
	s := &Stack[byte]{}
	s.Push('a')
	s.Push('b')
	Assert(t, s.Top() == 'b', "Stack Top is b")
	s.Pop()
	Assert(t, s.Top() == 'a', "Stack Top is a")
	s.Pop()
	Assert(t, !s.Pop(), "Pop Operation Failed")
	Assert(t, s.Len() == 0, "Stack is empty")
}
