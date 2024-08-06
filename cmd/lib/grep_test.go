package lib

import (
	"testing"
)

func TestMatchLiteral(t *testing.T) {
	regexp := NewByteIterator("a")
	text := NewByteIterator("apple")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching a Literal. 'a' in 'apple'")

	regexp.Reset()
	text = NewByteIterator("dog")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching a Literal. 'a' in 'dog'")
}

func TestMatchDigit(t *testing.T) {
	regexp := NewByteIterator(`\d`)
	text := NewByteIterator("apple 123")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching a Number. 1 in 'apple123'")

	regexp.Reset()
	text = NewByteIterator("dog")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching a Number. No Number in 'dog'")
}

func TestMatchAlphaNumeric(t *testing.T) {
	regexp := NewByteIterator(`\w`)
	text := NewByteIterator("alpha-num3ric")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching a AlphaNumber. 1 in 'alpha-num3ric'")

	regexp.Reset()
	text = NewByteIterator("$!?")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching a Number. No Number in '$!?'")
}

func TestPositiveGroup(t *testing.T) {
	regexp := NewByteIterator(`[abc]`)
	text := NewByteIterator("apple")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Positive Group [abc] in 'apple'")

	regexp.Reset()
	text = NewByteIterator("dog")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching Positive Group. [abc] in 'dog'")
}

func TestNegativeGroup(t *testing.T) {
	regexp := NewByteIterator(`[^abc]`)
	text := NewByteIterator("dog")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Negative Group [abc] in 'dog'")

	regexp.Reset()
	text = NewByteIterator("cab")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching Negative Group. [abc] in 'cab'")
}

func TestCombinationGroup(t *testing.T) {
	regexp := NewByteIterator(`\d\d\d apple`)
	text := NewByteIterator("123 apples")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Combinations Group '\\d\\d\\d apple' in '123 apples'")
}

func TestCombinationGroup1(t *testing.T) {
	regexp := NewByteIterator(`\d apple`)
	text := NewByteIterator("sally has 3 apples")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Combinations Group '\\d apple' in 'sally has 3 apples'")
}

// TODO: how this working
func TestCombinationGroup2(t *testing.T) {
	regexp := NewByteIterator(`\d apple`)
	text := NewByteIterator("sally has 300 apple")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Combinations Group '\\d apple' in 'sally has 300 ap'")
}

func TestMatchStar(t *testing.T) {
	regexp := NewByteIterator("c*")
	text := NewByteIterator("racer")

	result, _ := Match(text, regexp)
	Assert(t, result)
}

func TestMatchPlus(t *testing.T) {
	regexp := NewByteIterator("ca+t")
	text := NewByteIterator("caaat")

	result, _ := Match(text, regexp)
	Assert(t, result)
}

func TestAnchor(t *testing.T) {
	regexp := NewByteIterator("^log")
	text := NewByteIterator("logger")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching string anchor ^log with logger")

	regexp.Reset()
	text = NewByteIterator("slogger")
	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching string anchor ^log with slogger")
}

func TestLastAnchor(t *testing.T) {
	regexp := NewByteIterator("dog$")
	text := NewByteIterator("dog")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching string  dog$ with dog")
	regexp.Reset()
	text = NewByteIterator("dogs")
	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching string  dog$ with dogs")
}

func TestQuestion(t *testing.T) {
	regexp := NewByteIterator("dogs?")
	text := NewByteIterator("dog")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching string  dog? with dog")

	regexp.Reset()
	text = NewByteIterator("dogs")
	result, _ = Match(text, regexp)
	Assert(t, result, "Matching string dog? with dogs")

	regexp.Reset()
	text = NewByteIterator("cat")
	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching string dog? with cat")

}

func TestAlt(t *testing.T) {
	regexp := NewByteIterator("(cat|dog)")
	text := NewByteIterator("cat")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching string cat with (cat|dog)")

	regexp.Reset()
	text = NewByteIterator("dog")
	result, _ = Match(text, regexp)
	Assert(t, result, "Matching string dog with (cat|dog)")
}

func TestCaptureGroup(t *testing.T) {
	regexp := NewByteIterator(`(\w+) and dog (\d*)`)
	text := NewByteIterator("cat and dog 123")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching string '(\\w+) and dog' with cat and dog 123")
}
