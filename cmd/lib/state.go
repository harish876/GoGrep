package lib

// For tracking capture groups and other stateful operations like highlighting
type GrepState struct {
	Captures Stack[*ByteIterator]
}

func NewGrepState() *GrepState {
	return &GrepState{
		Captures: Stack[*ByteIterator]{},
	}
}

func (gs *GrepState) Reset() {
	gs.Captures = Stack[*ByteIterator]{}
}
