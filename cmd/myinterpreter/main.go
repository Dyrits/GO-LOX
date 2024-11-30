package main

import (
	"fmt"
	"os"
)

type Scanner struct{}

func (scanner *Scanner) Scan(content string) {
	const (
		LEFT_PAREN  rune = '('
		RIGHT_PAREN rune = ')'
		STAR        rune = '*'
		MINUS       rune = '-'
		PLUS        rune = '+'
		COMMA       rune = ','
		DOT         rune = '.'
		SEMICOLON   rune = ';'
		LEFT_BRACE  rune = '{'
		RIGHT_BRACE rune = '}'
	)

	runeNames := map[rune]string{
		LEFT_PAREN:  "LEFT_PAREN",
		RIGHT_PAREN: "RIGHT_PAREN",
		MINUS:       "MINUS",
		STAR:        "STAR",
		PLUS:        "PLUS",
		COMMA:       "COMMA",
		DOT:         "DOT",
		SEMICOLON:   "SEMICOLON",
		LEFT_BRACE:  "LEFT_BRACE",
		RIGHT_BRACE: "RIGHT_BRACE",
	}

	for _, character := range content {
		if name, exists := runeNames[character]; exists {
			fmt.Printf("%s %c null\n", name, character)
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
		scanner.Scan(string(content))
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
