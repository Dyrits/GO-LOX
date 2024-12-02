package main

import (
	"fmt"
	"os"
)

type Scanner struct{}

func (scanner *Scanner) Scan(content string) {
	tokens := map[rune]string{
		'(': "LEFT_PAREN",
		')': "RIGHT_PAREN",
		'-': "MINUS",
		'*': "STAR",
		'+': "PLUS",
		',': "COMMA",
		'.': "DOT",
		';': "SEMICOLON",
		'{': "LEFT_BRACE",
		'}': "RIGHT_BRACE",
		'=': "EQUAL",
		'!': "BANG",
		'>': "GREATER",
		'<': "LESS",
	}

	doubles := map[string]string{
		"==": "EQUAL_EQUAL",
		"!=": "BANG_EQUAL",
		">=": "GREATER_EQUAL",
		"<=": "LESS_EQUAL",
	}

	errors := map[rune]string{
		'#': "HASH",
		'$': "DOLLAR",
		'%': "PERCENT",
		'@': "AT",
	}

	invalid := false

	for index := 0; index < len(content); index++ {
		if (index + 1) < len(content) {
			if name, exists := doubles[content[index:index+2]]; exists {
				fmt.Println(fmt.Sprintf("%s %s null", name, content[index:index+2]))
				index++
				continue
			}
		}
		character := rune(content[index])
		if _, exists := errors[character]; exists {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("[line 1] Error: Unexpected character: %c", character))
			invalid = true
		}
		if name, exists := tokens[character]; exists {
			fmt.Println(fmt.Sprintf("%s %c null", name, character))
		}
	}
	fmt.Println("EOF  null")
	if invalid {
		os.Exit(65)
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
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
