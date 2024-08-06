package lib

// For tracking capture groups and other stateful operations like highlighting
type GrepState struct {
	Captures Stack[byte]
}

func NewGrepState() *GrepState {
	return &GrepState{
		Captures: Stack[byte]{},
	}
}
