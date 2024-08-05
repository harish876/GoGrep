package lib

import (
	"testing"
)

func TestIter(t *testing.T) {
	str := "harish_gokul"
	it := NewIterator(str)

	Assert(t, it.Get(19) == BUF_OUT_OF_RANGE)
	Assert(t, it.Get(0) == 'h')

	Assert(t, it.Advance(10).Peek() == 'u')
	Assert(t, it.Prev().Peek() == 'k')
	Assert(t, it.End() == 'l', "Last Char")
}
