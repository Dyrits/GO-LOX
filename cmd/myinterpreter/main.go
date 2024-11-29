package main

import (
	"fmt"
	"os"
)

type Scanner struct{}

func (scanner *Scanner) Parentheses(content string) {
	const (
		LEFT  rune = '('
		RIGHT rune = ')'
	)
	for _, character := range content {
		switch character {
		case LEFT:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT:
			fmt.Println("RIGHT_PAREN ) null")
		}
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(content) > 0 {
		scanner := &Scanner{}
		scanner.Parentheses(string(content))
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
