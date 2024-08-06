package lib

func Match(text *ByteIterator, regexp *ByteIterator) (bool, error) {
	if regexp.Get(0) == '^' {
		return MatchHere(text, regexp.Next()), nil
	}
	for {
		if MatchHere(text, regexp) {
			return true, nil
		}
		if !text.Next().HasNext() {
			return false, nil
		}
	}
}

func MatchHere(text *ByteIterator, regexp *ByteIterator) bool {
	if !regexp.HasNext() {
		return true
	} else if regexp.Peek(1) == '*' {
		return MatchStar(regexp.Peek(), text, regexp.Advance(2))
	} else if regexp.Peek(1) == '+' {
		return MatchPlus(regexp.Peek(), text, regexp.Advance(2))
	} else if regexp.Peek(1) == '?' {
		return MatchQuestion(regexp.Peek(), text, regexp.Advance(2))
	} else if regexp.Peek() == '$' {
		return !text.HasNext()
	} else if text.HasNext() && (regexp.Peek() == '.' || regexp.Peek() == text.Peek()) {
		return MatchHere(text.Next(), regexp.Next())
	} else if text.HasNext() && (regexp.Peek() == 0x5c && regexp.Peek(1) == 'd') {
		return MatchDigit(text, regexp)
	} else if text.HasNext() && (regexp.Peek() == 0x5c && regexp.Peek(1) == 'w') {
		return MatchAlphaNumeric(text, regexp)
	} else if text.HasNext() && (regexp.Peek() == '[' && regexp.Peek(1) != '^' && regexp.Find(']')) {
		return MatchPositiveGroup(text, regexp.Next())
	} else if text.HasNext() && (regexp.Peek() == '[' && regexp.Peek(1) == '^' && regexp.Find(']')) {
		return MatchNegativeGroup(text, regexp.Next())
	} else if text.HasNext() && (regexp.Peek() == '(' && regexp.Find(')')) {
		return MatchCaptureGroup(text, regexp.Next())
	}
	return false
}

func MatchCaptureGroup(text *ByteIterator, regexp *ByteIterator) bool {
	if regexp.Find('|') {
		return MatchOr(text, regexp)
	} else {
		var acc []byte
		for regexp.Peek() == ')' {
			acc = append(acc, regexp.Peek())
			regexp.Next()
		}
		captureRegex := NewByteIterator(acc)
		return MatchHere(text, captureRegex)
	}
}

func MatchOr(text *ByteIterator, regexp *ByteIterator) bool {
	var leftRegexp *ByteIterator
	var rightRegexp *ByteIterator
	var acc []byte
	for regexp.Peek() != ')' {
		if regexp.Peek() == '|' {
			leftRegexp = NewByteIterator(acc)
			acc = make([]byte, 0)
		} else {
			acc = append(acc, regexp.Peek())
		}
		regexp.Next()
	}
	rightRegexp = NewByteIterator(acc)
	text.Reset()
	leftMatch, _ := Match(text, leftRegexp)
	text.Reset()
	rightMatch, _ := Match(text, rightRegexp)
	return leftMatch || rightMatch
}

func MatchQuestion(char byte, text *ByteIterator, regexp *ByteIterator) bool {
	if MatchHere(text, regexp) {
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

func MatchAlphaNumeric(text *ByteIterator, regexp *ByteIterator) bool {
	if text.HasNext() && IsDigit(text.Peek()) || IsAlpha(text.Peek()) {
		regexp.Advance(2)
	}
	return MatchHere(text.Next(), regexp)
}

func MatchDigit(text *ByteIterator, regexp *ByteIterator) bool {
	if text.HasNext() && IsDigit(text.Peek()) {
		regexp.Advance(2)
	}

	return MatchHere(text.Next(), regexp)
}

func MatchStar(char byte, text *ByteIterator, regexp *ByteIterator) bool {
	for {
		if MatchHere(text, regexp) {
			return true
		}

		if !text.HasNext() || (text.Peek() != '.' && text.Peek() != char) {
			return false
		}
		text.Next()
	}
}

func MatchPlus(char byte, text *ByteIterator, regexp *ByteIterator) bool {
	for {
		if !text.HasNext() || (text.Peek() != '.' && text.Peek() != char) {
			return false
		}
		text.Next()
		if MatchHere(text, regexp) {
			return true
		}
	}
}
