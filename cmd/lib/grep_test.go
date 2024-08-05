package lib

import (
	"fmt"
	"testing"
)

func TestMatchLiteral(t *testing.T) {
	regexp := NewIterator("a")
	text := NewIterator("apple")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching a Literal. 'a' in 'apple'")

	regexp.Reset()
	text = NewIterator("dog")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching a Literal. 'a' in 'dog'")
}

func TestMatchDigit(t *testing.T) {
	regexp := NewIterator(`\d`)
	text := NewIterator("apple 123")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching a Number. 1 in 'apple123'")

	regexp.Reset()
	text = NewIterator("dog")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching a Number. No Number in 'dog'")
}

func TestMatchAlphaNumeric(t *testing.T) {
	regexp := NewIterator(`\w`)
	text := NewIterator("alpha-num3ric")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching a AlphaNumber. 1 in 'alpha-num3ric'")

	regexp.Reset()
	text = NewIterator("$!?")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching a Number. No Number in '$!?'")
}

func TestPositiveGroup(t *testing.T) {
	regexp := NewIterator(`[abc]`)
	text := NewIterator("apple")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Positive Group [abc] in 'apple'")

	regexp.Reset()
	text = NewIterator("dog")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching Positive Group. [abc] in 'dog'")
}

func TestNegativeGroup(t *testing.T) {
	regexp := NewIterator(`[^abc]`)
	text := NewIterator("dog")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Negative Group [abc] in 'dog'")

	regexp.Reset()
	text = NewIterator("cab")

	result, _ = Match(text, regexp)
	Assert(t, !result, "Matching Negative Group. [abc] in 'cab'")
}

func TestCombinationGroup(t *testing.T) {
	regexp := NewIterator(`\d\d\d apple`)
	text := NewIterator("123 apples")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Combinations Group '\\d\\d\\d apple' in '123 apples'")
}

func TestCombinationGroup1(t *testing.T) {
	regexp := NewIterator(`\d apple`)
	text := NewIterator("sally has 3 apples")

	result, _ := Match(text, regexp)
	Assert(t, result, "Matching Combinations Group '\\d apple' in 'sally has 3 apples'")
}

func TestMatchStar(t *testing.T) {
	regexp := NewIterator("c*")
	text := NewIterator("racer")

	result, _ := Match(text, regexp)
	fmt.Println(result)
}
