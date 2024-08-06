package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/codecrafters-io/grep-starter-go/cmd/lib"
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

	slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	text := lib.NewByteIterator(line)
	regexp := lib.NewByteIterator(pattern)
	gs := lib.NewGrepState()

	ok, err := lib.Match(text, regexp, gs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		fmt.Println("Failed to match. Match Result - ", ok)
		os.Exit(1)
	}

	fmt.Println("Match Result - ", ok)
}
