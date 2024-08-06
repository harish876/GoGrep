package lib

func Match(text *ByteIterator, regexp *ByteIterator, gs *GrepState) (bool, error) {
	if regexp.Get(0) == '^' {
		return MatchHere(text, regexp.Next(), gs), nil
	}
	for {
		if MatchHere(text, regexp, gs) {
			return true, nil
		}
		if !text.Next().HasNext() {
			return false, nil
		}
	}
}

func MatchHere(text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	if !regexp.HasNext() {
		return true
	} else if text.HasNext() && (regexp.Peek() == '(' && regexp.Find(')')) {
		AddCaptureGroup(text, regexp.Next(), gs)
		return MatchCaptureGroup(text, regexp.Next(), gs) && regexp.Peek() == BUF_OUT_OF_RANGE
	} else if regexp.Peek() == '*' {
		return MatchStar(regexp.Peek(-1), text, regexp.Advance(1), gs)
	} else if regexp.Peek() == '+' {
		return MatchPlus(regexp.Peek(-1), text, regexp.Advance(1), gs)
	} else if regexp.Peek(1) == '?' {
		return MatchQuestion(regexp.Peek(), text, regexp.Advance(1), gs)
	} else if regexp.Peek() == '$' {
		return !text.HasNext()
	} else if text.HasNext() && (regexp.Peek() == '.' || regexp.Peek() == text.Peek()) {
		return MatchHere(text.Next(), regexp.Next(), gs)
	} else if text.HasNext() && (regexp.Peek() == 0x5c && regexp.Peek(1) == 'd') {
		return MatchDigit(text, regexp, gs)
	} else if text.HasNext() && (regexp.Peek() == 0x5c && regexp.Peek(1) == 'w') {
		return MatchAlphaNumeric(text, regexp, gs)
	} else if text.HasNext() && (regexp.Peek() == 0x5c && IsDigit(regexp.Peek(1)) && gs.Captures.Len() >= 1) {
		capture := gs.Captures.Top()
		//gs.Captures.Pop()
		result := MatchHere(text, capture, gs)
		return result
	} else if text.HasNext() && (regexp.Peek() == '[' && regexp.Peek(1) != '^' && regexp.Find(']')) {
		return MatchPositiveGroup(text, regexp.Next())
	} else if text.HasNext() && (regexp.Peek() == '[' && regexp.Peek(1) == '^' && regexp.Find(']')) {
		return MatchNegativeGroup(text, regexp.Next())
	}
	return false
}

func AddCaptureGroup(text *ByteIterator, regexp *ByteIterator, gs *GrepState) {
	var acc []byte
	for regexp.Peek() != ')' {
		acc = append(acc, regexp.Peek())
		regexp.Next()
	}
	gs.Captures.Push(NewByteIterator(acc))
}

func MatchCaptureGroup(text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	if gs.Captures.Len() == 0 {
		return false
	}

	capture := gs.Captures.Top()
	//gs.Captures.Pop()

	if capture.Find('|') {
		return MatchOr(text, capture, gs) // MatchOr should pop the game state
	} else {
		var acc []byte
		for capture.Peek() == ')' {
			acc = append(acc, capture.Peek())
			capture.Next()
		}
		captureRegex := NewByteIterator(acc)
		return MatchHere(text, captureRegex, gs)
	}
}

func MatchOr(text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	var leftRegexp *ByteIterator
	var rightRegexp *ByteIterator
	var acc []byte

	for regexp.HasNext() {
		if regexp.Peek() == '|' {
			leftRegexp = NewByteIterator(acc)
			acc = make([]byte, 0)
		} else {
			acc = append(acc, regexp.Peek())
		}
		regexp.Next()
	}
	rightRegexp = NewByteIterator(acc)

	leftMatch, _ := Match(text, leftRegexp, gs)
	text.Reset()
	rightMatch, _ := Match(text, rightRegexp, gs)
	return leftMatch || rightMatch
}

func MatchQuestion(char byte, text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	if MatchHere(text, regexp, gs) {
		return true
	}
	if text.HasNext() || (text.Peek() != '.' && text.Peek() != char) {
		return false
	}
	return true
}

func MatchPositiveGroup(text *ByteIterator, regexp *ByteIterator) bool {
	charSet := make(map[byte]bool)
	for regexp.Peek() != ']' {
		charSet[regexp.Peek()] = true
		regexp.Next()
	}
	for text.HasNext() {
		if _, ok := charSet[text.Peek()]; ok {
			return true
		}
		text.Next()
	}
	return false
}

func MatchNegativeGroup(text *ByteIterator, regexp *ByteIterator) bool {
	charSet := make(map[byte]bool)
	for regexp.Peek() != ']' {
		charSet[regexp.Peek()] = true
		regexp.Next()
	}
	for text.HasNext() {
		if _, ok := charSet[text.Peek()]; ok {
			return false
		}
		text.Next()
	}
	return true
}

func MatchAlphaNumeric(text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	if text.HasNext() && IsDigit(text.Peek()) || IsAlpha(text.Peek()) {
		regexp.Advance(2)
	}
	return MatchHere(text.Next(), regexp, gs)
}

func MatchDigit(text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	if text.HasNext() && IsDigit(text.Peek()) {
		regexp.Advance(2)
	}

	return MatchHere(text.Next(), regexp, gs)
}

func MatchStar(char byte, text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	for {
		if MatchHere(text, regexp, gs) {
			return true
		}

		if !text.HasNext() || (text.Peek() != '.' && text.Peek() != char) {
			return false
		}
		text.Next()
	}
}

func MatchPlus(char byte, text *ByteIterator, regexp *ByteIterator, gs *GrepState) bool {
	for {
		if !text.HasNext() || (text.Peek() != '.' && text.Peek() != char) {
			return false
		}
		text.Next()
		if MatchHere(text, regexp, gs) {
			return true
		}
	}
}
