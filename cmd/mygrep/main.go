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
		return ok, nil
	case pattern == `\w`:
		ok = matchDigitOrChar(line)
		return ok, nil
	case len(pattern) >= 2 && pattern[0] == '[' && pattern[len(pattern)-1] == ']':
		charsToMatch := make(map[byte]bool, 0)
		for i := 1; i < len(pattern); i++ {
			charsToMatch[pattern[i]] = true
		}
		ok = matchCharSet(line, charsToMatch)
		return ok, nil
	default:
		ok = bytes.ContainsAny(line, pattern)
		return ok, nil
	}
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

func matchCharSet(line []byte, charSet map[byte]bool) bool {
	for _, char := range line {
		if _, ok := charSet[char]; ok {
			return true
		}
	}
	return false
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
