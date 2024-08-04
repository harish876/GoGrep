package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}
}

func matchLine(line []byte, pattern string) (bool, error) {
	var ok bool
	switch {
	case pattern == `\d`:
		ok = matchDigit(line)
	case pattern == `\w`:
		ok = matchDigitOrChar(line)
	case isPositiveCharGroup(pattern):
		ok = matchCharSet(line, pattern)
	case isNegativeCharGroup(pattern):
		ok = matchNoneInCharSet(line, pattern)
	default:
		ok = bytes.ContainsAny(line, pattern)
	}

	return ok, nil

}

func matchDigit(line []byte) bool {
	for _, char := range line {
		if isDigit(char) {
			return true
		}
	}
	return false
}

func matchChar(line []byte) bool {
	for _, char := range line {
		if isAlpha(char) {
			return true
		}
	}
	return false
}

func matchDigitOrChar(line []byte) bool {
	for _, char := range line {
		if isAlpha(char) || isDigit(char) {
			return true
		}
	}
	return false
}

// any
func matchCharSet(line []byte, pattern string) bool {
	charsToMatch := make(map[byte]bool, 0)
	for i := 1; i < len(pattern); i++ {
		charsToMatch[pattern[i]] = true
	}
	for _, char := range line {
		if _, ok := charsToMatch[char]; ok {
			return true
		}
	}
	return false
}

// none
func matchNoneInCharSet(line []byte, pattern string) bool {
	charsToMatch := make(map[byte]bool, 0)
	for i := 1; i < len(pattern); i++ {
		charsToMatch[pattern[i]] = true
	}
	for _, char := range line {
		if _, ok := charsToMatch[char]; ok {
			return false
		}
	}
	return true
}
func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isPositiveCharGroup(pattern string) bool {
	n := len(pattern)
	if n < 2 {
		return false
	}
	return pattern[0] == '[' && pattern[n-1] == ']' && pattern[1] != '^'
}

func isNegativeCharGroup(pattern string) bool {
	n := len(pattern)
	if n < 2 {
		return false
	}
	return pattern[0] == '[' && pattern[n-1] == ']' && pattern[1] == '^'
}
