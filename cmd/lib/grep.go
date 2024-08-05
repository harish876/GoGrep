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
	}
	if regexp.Peek(1) == '*' {
		return MatchStar(regexp.Peek(), text, regexp.Advance(2))
	}
	if regexp.Peek(1) == '+' {
		return MatchPlus(regexp.Peek(), text, regexp.Advance(2))
	}
	if regexp.Peek() == '$' {
		return !text.HasNext()
	}
	if text.HasNext() && (regexp.Peek() == '.' || regexp.Peek() == text.Peek()) {
		return MatchHere(text.Next(), regexp.Next())
	}
	if text.HasNext() && (regexp.Peek() == 0x5c && regexp.Peek(1) == 'd') {
		return MatchDigit(text, regexp)
	}
	if text.HasNext() && (regexp.Peek() == 0x5c && regexp.Peek(1) == 'w') {
		return MatchAlphaNumeric(text, regexp)
	}
	if text.HasNext() && (regexp.Peek() == '[' && regexp.End() == ']' && regexp.Peek(1) != '^') {
		return MatchPositiveGroup(text, regexp.Next())
	}
	if text.HasNext() && (regexp.Peek() == '[' && regexp.End() == ']' && regexp.Peek(1) == '^') {
		return MatchNegativeGroup(text, regexp.Next())
	}
	return false
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
